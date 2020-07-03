package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/ext/transforms/doFn"

	csvsnk "github.com/raralabs/hike/plugins/canal/sinks"
	csvsrc "github.com/raralabs/hike/plugins/canal/sources"
)

const TmpPath = "./tmp/"

func main() {
	if err := os.MkdirAll(TmpPath, os.ModePerm); err != nil {
		log.Panic(err)
	}

	p := pipeline.NewPipeline(1)
	opts := pipeline.DefaultProcessorOptions

	file, err := os.Open(TmpPath + "records.csv")
	if err != nil {
		log.Panic(err)
	}

	src := p.AddSource("CSV Reader")
	sp := src.AddProcessor(opts, csvsrc.NewCsvReader(file, -1))

	delay := p.AddTransform("Delay")
	del := delay.AddProcessor(opts, doFn.DelayFunction(100*time.Millisecond), "path1")

	snk := p.AddSink("CSV Writer")
	snk.AddProcessor(opts, csvsnk.NewCsvWriter(os.Stdout), "sink")

	delay.ReceiveFrom("path1", sp)
	snk.ReceiveFrom("sink", del)

	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	p.Validate()
	p.Start(c, cancel)
}
