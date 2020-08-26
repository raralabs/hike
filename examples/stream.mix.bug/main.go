package main

import (
	"context"
	"github.com/Knetic/govaluate"
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/ext/sources"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/hike/plugins/canal/sinks"
	"log"
	"os"
	"time"
)

func main() {

	p := pipeline.NewPipeline(1)

	opts := pipeline.DefaultProcessorOptions
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

	src := p.AddSource("Fake Source")
	sp := src.AddProcessor(opts, sources.NewFaker(10, nil))

	tr1 := p.AddTransform("Filter Stage")
	pr1 := tr1.AddProcessor(opts, doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
		return true, nil
	}, doneFunc), "path1")

	tr2 := p.AddTransform("Enrich Stage")
	expr, err := govaluate.NewEvaluableExpression("age * 2")
	if err != nil {
		log.Println(err)
		return
	}
	pr2 := tr2.AddProcessor(opts, doFn.EnrichFunction("twice_age", expr, doneFunc), "path2")

	snk1 := p.AddSink("Filter Sink")
	snk1.AddProcessor(opts, sinks.NewPrettyPrinter(os.Stdout, 10), "path3")
	//snk1.AddProcessor(opts, canalSnk.NewBlackholeSink(), "path3")

	snk2 := p.AddSink("Enrich Sink")
	snk2.AddProcessor(opts, sinks.NewPrettyPrinter(os.Stdout, 5), "path4")
	//snk2.AddProcessor(opts, canalSnk.NewBlackholeSink(), "path4")

	tr1.ReceiveFrom("path1", sp)
	tr2.ReceiveFrom("path2", sp)
	snk1.ReceiveFrom("path3", pr1)
	snk2.ReceiveFrom("path4", pr2)

	p.Validate()
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	p.Start(c, cancel)
}
