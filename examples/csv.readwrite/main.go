package main

import (
	"context"
	"time"

	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/transforms"
	csvsnk "github.com/raralabs/hike/plugins/canal/sinks"
	csvsrc "github.com/raralabs/hike/plugins/canal/sources"
)

func main() {
	p := pipeline.NewPipeline(1)

	src := p.AddSource("CSV Reader")
	sp := src.AddProcessor(csvsrc.NewCsvReader("./tmp/records.csv", -1))

	delay := p.AddTransform("Delay")
	del := delay.AddProcessor(transforms.DelayFunction(100*time.Millisecond), "path1")

	snk := p.AddSink("CSV Writer")
	snk.AddProcessor(csvsnk.NewCsvWriter("./tmp/newRecords.csv"), "sink")

	delay.ReceiveFrom("path1", sp)
	snk.ReceiveFrom("sink", del)

	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	p.Validate()
	p.Start(c, cancel)
}
