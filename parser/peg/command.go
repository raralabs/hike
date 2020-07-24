package peg

import (
	"fmt"
	"log"
	"os"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/core/transforms/agg"
	"github.com/raralabs/canal/ext/transforms/aggregates"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/canal/utils/cast"

	"github.com/raralabs/hike/plugins/canal/sinks"
	"github.com/raralabs/hike/plugins/canal/sources"
)

var availablePicks = []string{
	"first",
	"random",
	"last",
}

func strIn(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

type command struct {
}

func (c *command) Init() {
	// Build the grammar, generate codes and stuffs in here
}

func (c *command) Build(id uint32, cmd string) (startFunc func(), ppln *pipeline.Pipeline, closeFunc func()) {
	stages, err := Parse("", []byte(cmd))

	if err != nil {
		log.Panic(err)
	}

	p := pipeline.NewPipeline(id)
	opts := pipeline.DefaultProcessorOptions

	var aggregators []*agg.Aggregator
	var openedFiles []*os.File
	var lastProc pipeline.IProcessor

	useDefaultSink := true

	stgs := cast.ToIfaceSlice(stages)
	for i, stg := range stgs {
		pathName := fmt.Sprintf("path%v", i)
		routeParam := pipeline.MsgRouteParam(pathName)
		stageName := fmt.Sprintf("Stage%v", i)

		switch s := stg.(type) {
		case FileJob:
			// FileJob can act as both a source or sink depending upon it's placement
			fileName := s.Filename

			// Act as source if at the beginning
			if i == 0 {
				f, err := os.Open(fileName)
				if err != nil {
					log.Panic(err)
				}
				log.Printf("[INFO] %s file opened as source.", fileName)
				src := p.AddSource("File Source")
				sp := src.AddProcessor(opts, sources.NewCsvReader(f, -1))

				lastProc = sp
				openedFiles = append(openedFiles, f)
			} else if i == len(stgs)-1 { // Act as sink if at the end
				f, err := os.Create(fileName)
				if err != nil {
					log.Panic(err)
				}
				log.Printf("[INFO] %s file created as sink.", fileName)

				snk := p.AddSink("File Sink")
				snk.AddProcessor(opts, sinks.NewCsvWriter(f), routeParam)

				snk.ReceiveFrom(routeParam, lastProc)
				openedFiles = append(openedFiles, f)
				useDefaultSink = false
			}

		case DoNodeJob:
			stg := p.AddTransform(stageName)
			var proc pipeline.IProcessor

			doneFunc := func(m message.Msg) bool {
				content := m.Content()

				if content != nil {
					if v, ok := content.Get("eof"); ok {
						if v.Val == true {
							return true
						}
					}
				}
				return false
			}

			switch doJob := s.Function.(type) {
			case Filter:
				proc = stg.AddProcessor(opts, doFn.FilterFunction(func(m message.Msg) (bool, bool, error) {

					content := m.Content()
					if content == nil {
						return true, false, nil
					}

					if v, ok := content.Get("eof"); ok {
						if v.Val == true {
							return true, true, nil
						}
					}
					match, err := doJob(m.Values())
					return match, false, err
				}), routeParam)

			case Select:
				fields := doJob.Fields
				proc = stg.AddProcessor(opts, doFn.SelectFunction(fields, doneFunc), routeParam)

			case Pick:
				dsc := doJob.Desc
				if !strIn(availablePicks, dsc) {
					log.Panicf("Can't pick by: %s", dsc)
				}
				proc = stg.AddProcessor(opts, doFn.PickFunction(dsc, doJob.Num, doneFunc), routeParam)

			case Sort:
				fld := doJob.Field
				proc = stg.AddProcessor(opts, doFn.SortFunction(fld, doneFunc), routeParam)

			case Batch:
				proc = stg.AddProcessor(opts, doFn.BatchAgg(doneFunc), routeParam)
			}

			stg.ReceiveFrom(routeParam, lastProc)
			lastProc = proc

		case AggNodeJob:
			aggFuncs := s.Functions
			var aggs []agg.IAggFuncTemplate

			after := func(m message.Msg, proc pipeline.IProcessorForExecutor, content, pContent []*message.OrderedContent) {

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
				for i := range content {
					proc.Result(m, content[i], pContent[i])
				}
			}

			for _, ags := range aggFuncs {
				switch ag := ags.(type) {
				case Count:
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

				case Max:
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

				case Min:
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

				case Avg:
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

				case Sum:
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

				case Variance:
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

				case DistinctCount:
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

				case Mode:
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

				case Quantile:
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

				case Covariance:
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

				case Correlation:
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

				case PValue:
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

			aggregators = append(aggregators, aggregator)

			stg := p.AddTransform(stageName)
			proc := stg.AddProcessor(opts, aggregator.Function(), routeParam)

			stg.ReceiveFrom(routeParam, lastProc)
			lastProc = proc

		case SinkJob:
			if i != len(stgs)-1 {
				log.Panic("Sink Job not at the End")
			}

			switch s.Type {
			case "stdout":
				useDefaultSink = true

			default:
				switch snkJob := s.Args.(type) {
				case Plot:
					snk := p.AddSink("Plot Sink")
					switch w := snkJob.Widget.(type) {
					case BarPlot:
						snk.AddProcessor(opts, sinks.NewBarPlot(w.Title, w.XField, w.YField, w.BarWidth, w.BarGap), routeParam)
					}
					snk.ReceiveFrom(routeParam, lastProc)
				}
			}
		}

		if i == len(stgs)-1 && useDefaultSink {
			defaultSink(p, routeParam, lastProc)
		}
	}
	p.Validate()

	starter := func() {
		//if aggregators != nil && len(aggregators) != 0 {
		//	for _, ag := range aggregators {
		//		ag.Start()
		//	}
		//}
	}

	closer := func() {
		if openedFiles != nil && len(openedFiles) != 0 {
			for _, file := range openedFiles {
				if err := file.Close(); err != nil {
					log.Panic(err)
				}
				log.Printf("[INFO] %s file closed", file.Name())
			}
		}
	}

	return starter, p, closer
}

func NewPegCmd() *command {
	return &command{}
}

func defaultSink(p *pipeline.Pipeline, routeParam pipeline.MsgRouteParam, lastProc pipeline.IProcessor) {
	opts := pipeline.DefaultProcessorOptions

	snk := p.AddSink("Stdout Sink")
	snk.AddProcessor(opts, sinks.NewPrettyPrinter(os.Stdout, 10), routeParam)

	snk.ReceiveFrom(routeParam, lastProc)
}
