package main

import (

"github.com/raralabs/canal/core/message"
"github.com/raralabs/canal/core/message/content"
"github.com/raralabs/canal/core/pipeline"
"github.com/raralabs/canal/core/transforms/agg"
canalSnk "github.com/raralabs/canal/ext/sinks"
canalSrc "github.com/raralabs/canal/ext/sources"
"github.com/raralabs/canal/ext/transforms/aggregates"
"github.com/raralabs/canal/ext/transforms/doFn"
"github.com/raralabs/canal/utils/cast"
"github.com/raralabs/hike/parser/at"
"github.com/raralabs/hike/parser/peg"
"github.com/raralabs/hike/plugins/canal/sinks"
"github.com/raralabs/hike/plugins/canal/sources"
	"hike/hike/parser/newPeg"
	"log"
"os"

"sync"

)

type atBuilder struct {
	streamFromMu, streamToMu *sync.Mutex
	streamFrom               map[string][]at.Node
	streamTo                 map[string]*node
}

func newATBuilder() *atBuilder {
	return &atBuilder{
		streamFromMu: &sync.Mutex{},
		streamToMu:   &sync.Mutex{},
		streamFrom:   make(map[string][]at.Node),
		streamTo:     make(map[string]*node),
	}
}

func (p *atBuilder) BuildAT(cmds []interface{}) at.AT{
	startId := int64(1)
	var srcs []at.Node
	for _,command := range cmds{
		var src at.Node
		src, startId = p.buildSinglePipeline(startId,cast.ToIfaceSlice(command))
		if src != nil{
			srcs = append(srcs,src)
		}
	}
	p.streamToMu.Lock()
	p.streamFromMu.Lock()
	defer p.streamToMu.Unlock()
	defer p.streamFromMu.Unlock()

	for s := range p.streamTo {
		if tos, ok := p.streamFrom[s]; ok {
			p.streamTo[s].toNodes = append(p.streamTo[s].toNodes, tos...)
		}
	}

	absTree := &tree{sources: srcs}

	return absTree


}


func (p *atBuilder) buildSinglePipeline(Id int64,command []interface{})(at.Node, int64) {
	streamToName := ""
	streamFromName := ""
	var lastNode *node
	var firstNode *node
	for i, eachStage := range command {
		switch eachStage.(type){
		case newPeg.Into:
			streamToName = eachStage.(newPeg.Into).StreamTo
		case newPeg.BranchJob:
			streamFromName = eachStage.(newPeg.BranchJob).BranchName
			exec := getExecutor(eachStage)
			if exec == nil {
				log.Panic(eachStage)
			}
			n := &node{
				id:    Id,
				exec:  exec,
				added: true,
			}
			Id++
			p.streamFromMu.Lock()
			p.streamFrom[streamFromName] = append(p.streamFrom[streamFromName], n)
			p.streamFromMu.Unlock()
		default:
			exec := getExecutor(eachStage)
			if exec == nil {
				log.Panic(eachStage)
			}
			n := &node{
				id:    Id,
				exec:  exec,
				added: true,
			}
			Id++

			if i > 0 && lastNode != nil {
				lastNode.toNodes = append(lastNode.toNodes, n)
			}else{
				firstNode = n
			}



			// If last node, and streamTo exists, add to streamTo
			if i == len(command)-1 && streamToName != "" {
				p.streamToMu.Lock()
				var keys []string
				for k := range p.streamTo {
					keys = append(keys, k)
				}
				if _, ok := p.streamTo[streamToName]; ok {
					p.streamToMu.Unlock()
					log.Panic("Can't have two different pipeline streaming to single stream")
				}
				p.streamTo[streamToName] = n
				p.streamToMu.Unlock()
			}
			lastNode = n
		}

		}
		if firstNode == nil {
			return nil, Id
		}
		return firstNode, Id


}


func getExecutor(stg interface{}) pipeline.Executor {

	var exec pipeline.Executor
	switch s := stg.(type) {
	case peg.FileJob:
		// FileJob can act as both a source or sink depending upon it's placement
		fileName := s.Filename

		f, err := os.Open(fileName)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("[INFO] %s file opened as source.", fileName)
		exec = sources.NewCsvReader(f, -1)
		exec.SetName("file")

	case peg.FakeJob:
		// FileJob can act as both a source or sink depending upon it's placement
		numData := s.NumData

		log.Printf("[INFO] Generating %d fake data", numData)
		exec = canalSrc.NewFaker(numData, nil)
		exec.SetName("fake")

	case peg.DoNodeJob:

		doneFunc := func(m message.Msg) bool {
			cont := m.Content()

			if cont != nil {
				if v, ok := cont.Get("eof"); ok {
					if v.Val == true {
						return true
					}
				}
			}
			return false
		}

		switch doJob := s.Function.(type) {
		case peg.Filter:
			exec = doFn.FilterFunction(doJob, doneFunc)
			exec.SetName("filter")

		case peg.Select:
			fields := doJob.Fields
			exec = doFn.SelectFunction(fields, doneFunc)
			exec.SetName("select")

		case peg.Pick:
			dsc := doJob.Desc
			//if !strIn(availablePicks, dsc) {
			//	log.Panicf("Can't pick by: %s", dsc)
			//}
			exec = doFn.PickFunction(dsc, doJob.Num, doneFunc)
			exec.SetName("pick")

		case peg.Sort:
			fld := doJob.Field
			exec = doFn.SortFunction(fld, doneFunc)
			exec.SetName("sort")

		case peg.Batch:
			exec = doFn.BatchAgg(doneFunc)
			exec.SetName("batch")

		case peg.Enrich:
			if expr := doJob.Expr; expr != nil {
				exec = doFn.EnrichFunction(doJob.Field, expr, doneFunc)
				exec.SetName("enrich")
			}
		}

	case peg.AggNodeJob:
		aggFuncs := s.Functions
		var aggs []agg.IAggFuncTemplate

		after := func(m message.Msg, proc pipeline.IProcessorForExecutor, cntnt, pContent []content.IContent) {

			contents := m.Content()
			if contents != nil {
				if v, ok := contents.Get("eof"); ok {
					if v.Val == true {
						proc.Result(m, contents, nil)
						proc.Done()
						return
					}
				}
			}
			for i := range cntnt {
				proc.Result(m, cntnt[i], pContent[i])
			}
		}

		for _, ags := range aggFuncs {
			switch ag := ags.(type) {
			case peg.Count:
				cnt := aggregates.NewCount(ag.Alias, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, cnt)

			case peg.Max:
				mx := aggregates.NewMax(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panicf("Max Filter Error: %v", err)
					}
					return match
				})
				aggs = append(aggs, mx)

			case peg.Min:
				mn := aggregates.NewMin(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, mn)

			case peg.Avg:
				avg := aggregates.NewAvg(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, avg)

			case peg.Sum:
				sum := aggregates.NewSum(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, sum)

			case peg.Variance:
				variance := aggregates.NewVariance(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, variance)

			case peg.DistinctCount:
				variance := aggregates.NewDCount(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, variance)

			case peg.Mode:
				variance := aggregates.NewMode(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, variance)

			case peg.Quantile:
				quantile := aggregates.NewQuantile(ag.Alias, ag.Field, ag.Qth,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, quantile)

			case peg.Covariance:
				cov := aggregates.NewCovariance(ag.Alias, ag.Field1, ag.Field2,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, cov)

			case peg.Correlation:
				corr := aggregates.NewCorrelation(ag.Alias, ag.Field1, ag.Field2,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, corr)

			case peg.PValue:
				pval := aggregates.NewPValue(ag.Alias, ag.Field1, ag.Field2,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, pval)
			}
		}

		aggregator := agg.NewAggregator(aggs, after, s.GroupBy...)

		exec = aggregator.Function()
		exec.SetName("agg")

	case peg.SinkJob:

		switch s.Type {
		case "stdout":

			exec = sinks.NewPrettyPrinter(os.Stdout, 10, s.Header...)
			exec.SetName("stdout")

		case "blackhole":
			exec = canalSnk.NewBlackholeSink()
			exec.SetName("blackhole")

		default:
			switch snkJob := s.Args.(type) {
			case peg.Plot:
				switch w := snkJob.Widget.(type) {
				case peg.BarPlot:
					exec = sinks.NewBarPlot(w.Title, w.XField, w.YField, w.BarWidth, w.BarGap)
					exec.SetName("plot")
				}
			}
		}
	}
	if exec == nil {
		return nil
	}

	return exec
}
