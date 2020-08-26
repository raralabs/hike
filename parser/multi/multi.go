package multi

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
	"log"
	"os"
	"strings"
	"sync"
	"unicode"
)

type node struct {
	id       int64
	executor pipeline.Executor
	added    bool
	toNodes  []at.Node
}

func (n *node) ID() int64 {
	return n.id
}

func (n *node) Executor() pipeline.Executor {
	return n.executor
}

func (n *node) Add() {
	n.added = true
}

func (n *node) IsAdded() bool {
	return n.added
}

func (n *node) ToNodes() []at.Node {
	return n.toNodes
}

type tree struct {
	sources []at.Node
}

func (t *tree) Sources() []at.Node {
	return t.sources
}

var (
	streamFromMu, streamToMu *sync.Mutex
	streamFrom               map[string][]at.Node
	streamTo                 map[string]*node
)

func init() {
	streamFromMu = &sync.Mutex{}
	streamToMu = &sync.Mutex{}
	streamFrom = make(map[string][]at.Node)
	streamTo = make(map[string]*node)
}

func Build(cmds ...string) at.AT {

	// Remove ";" from the end of each command if present.
	// Remove "into xxx" from the end.
	//		-> Add the output coming to "xxx" to the new input if the stream starts with "xxx".
	//		Eg: fake() | select(age) into s1
	//			s1 | map twice_age = 2*age
	//		-> In such cases the output of select must go to the map.

	startId := int64(1)
	var srcs []at.Node

	for _, c := range cmds {
		var src at.Node
		src, startId = buildSinglePipe(startId, c)
		if src != nil {
			srcs = append(srcs, src)
		}
	}

	streamToMu.Lock()
	streamFromMu.Lock()
	defer streamToMu.Unlock()
	defer streamFromMu.Unlock()

	for s := range streamTo {
		if tos, ok := streamFrom[s]; ok {
			streamTo[s].toNodes = append(streamTo[s].toNodes, tos...)
		}
	}
	absTree := &tree{sources: srcs}

	return absTree
}

func buildSinglePipe(startId int64, cmd string) (at.Node, int64) {
	cmd = strings.TrimSpace(cmd)
	if cmd[len(cmd)-1] == ';' {
		cmd = cmd[:len(cmd)-1]
	}

	// Check if the output goes to some stream

	cmds := strings.Fields(cmd)
	streamToName := ""
	streamFromName := ""

	if len(cmds) > 1 {
		if cmds[len(cmds)-2] == "into" {
			streamToName = cmds[len(cmds)-1]
			cmds = cmds[:len(cmds)-2]
			cmd = strings.Join(cmds, " ")
		}
	}

	if isVariable(cmds[0]) {
		// This means the pipeline starts with stream.
		streamFromName = cmds[0]
		cmds := strings.Split(cmd, "|")
		cmd = strings.Join(cmds[1:], "|")
	}
	cmd = strings.TrimSpace(cmd)
	//fmt.Println(cmd)

	// cmd is now ready to be parsed by peg
	stages, err := peg.Parse("", []byte(cmd))
	if err != nil {
		log.Panic(err)
	}
	stgs := cast.ToIfaceSlice(stages)

	var lastNode *node
	var firstNode *node

	for i, stg := range stgs {
		exec := getExecutor(stg)
		if exec == nil {
			log.Panic(stg)
		}
		n := &node{
			id:       startId,
			executor: exec,
			added:    true,
		}
		startId++

		if i > 0 && lastNode != nil {
			lastNode.toNodes = append(lastNode.toNodes, n)
		}

		// If first node, and streamFromName exists, add to streamFrom
		if i == 0 {
			if streamFromName != "" {
				streamFromMu.Lock()
				streamFrom[streamFromName] = append(streamFrom[streamFromName], n)
				streamFromMu.Unlock()
			} else {
				// No streamFromName implies this node must be a source
				firstNode = n
			}
		}

		// If last node, and streamTo exists, add to streamTo
		if i == len(stgs)-1 && streamToName != "" {
			streamToMu.Lock()
			if _, ok := streamTo[streamToName]; ok {
				streamToMu.Unlock()
				log.Panic("Can't have two different pipeline streaming to single stream")
			}
			streamTo[streamToName] = n
			streamToMu.Unlock()
		}

		lastNode = n
	}
	if firstNode == nil {
		return nil, startId
	}

	return firstNode, startId
}

func isVariable(v string) bool {
	if !(unicode.IsLetter(rune(v[0])) || v[0] == '_') {
		return false
	}

	for _, l := range v {
		if !(unicode.IsLetter(l) || l == '_' || unicode.IsDigit(l)) {
			return false
		}
	}
	return true
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
