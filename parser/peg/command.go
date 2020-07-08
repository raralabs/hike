package peg

import (
	"fmt"
	"log"
	"os"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/core/transforms/agg"
	"github.com/raralabs/canal/core/transforms/event/poll"
	"github.com/raralabs/canal/ext/transforms/aggregates"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/canal/utils/cast"

	"github.com/raralabs/hike/plugins/canal/sinks"
	"github.com/raralabs/hike/plugins/canal/sources"
)

type Command struct {
}

func (c *Command) Init() {
	// Build the grammar, generate codes and stuffs in here
}

func (c *Command) Build(id uint32, cmd string) (startFunc func(), ppln *pipeline.Pipeline, closeFunc func()) {
	stages, err := Parse("", []byte(cmd))

	if err != nil {
		log.Panic(err)
	}

	p := pipeline.NewPipeline(id)
	opts := pipeline.DefaultProcessorOptions

	var aggregators []*agg.Aggregator
	var openedFiles []*os.File
	var lastProc pipeline.IProcessor

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
			if i == 0 && len(stgs) > 1 {
				f, err := os.Open(fileName)
				if err != nil {
					log.Panic(err)
				}
				log.Printf("[INFO] %s file opened as source.", fileName)
				src := p.AddSource("File Source")
				sp := src.AddProcessor(opts, sources.NewCsvReader(f, -1))

				lastProc = sp
				openedFiles = append(openedFiles, f)
			}

			// Act as sink if at the end
			if i == len(stgs)-1 {
				f, err := os.Create(fileName)
				if err != nil {
					log.Panic(err)
				}
				log.Printf("[INFO] %s file created as sink.", fileName)

				snk := p.AddSink("File Sink")
				snk.AddProcessor(opts, sinks.NewCsvWriter(f), routeParam)

				snk.ReceiveFrom(routeParam, lastProc)
				openedFiles = append(openedFiles, f)
			}

		case DoNodeJob:
			switch doJob := s.Function.(type) {
			case Filter:
				stg := p.AddTransform(stageName)
				proc := stg.AddProcessor(opts, doFn.FilterFunction(func(m message.Msg) (bool, bool, error) {
					content := m.Content()

					if v, ok := content.Get("eof"); ok {
						if v.Val == true {
							return true, true, nil
						}
					}
					match, err := doJob(m.Values())
					return match, false, err
				}), routeParam)

				stg.ReceiveFrom(routeParam, lastProc)
				lastProc = proc
			}

		case AggNodeJob:
			aggFuncs := s.Functions
			var aggs []agg.IAggregator
			for _, ags := range aggFuncs {
				switch ag := ags.(type) {
				case Count:
					cnt := aggregates.NewCount(ag.Alias, func(m map[string]interface{}) bool {
						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
					aggs = append(aggs, cnt)

				case Max:
					mx := aggregates.NewMax(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
						match, err := ag.Filter(m)
						if err != nil {
							log.Panicf("Max Filter Error: %v", err)
						}
						return match
					})
					aggs = append(aggs, mx)

				case Min:
					mn := aggregates.NewMin(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
					aggs = append(aggs, mn)
				}
			}

			after := func(m message.Msg, proc pipeline.IProcessorForExecutor, msgs []*message.OrderedContent) bool {
				content := m.Content()

				if v, ok := content.Get("eof"); ok {
					if v.Val == true {
						for _, msg := range msgs {
							proc.Result(m, msg)
						}

						proc.Result(m, content)
						return true
					}
				}
				return false
			}

			aggregator := agg.NewAggregator(poll.NewFilterEvent(func(m map[string]interface{}) bool {
				return true
			}), aggs, after, s.GroupBy...)

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
				snk := p.AddSink("Stdout Sink")
				snk.AddProcessor(opts, sinks.NewCsvWriter(os.Stdout), routeParam)

				snk.ReceiveFrom(routeParam, lastProc)
			}
		}
	}
	p.Validate()

	starter := func() {
		if aggregators != nil && len(aggregators) != 0 {
			for _, ag := range aggregators {
				ag.Start()
			}
		}
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

func NewPegCmd() *Command {
	return &Command{}
}
