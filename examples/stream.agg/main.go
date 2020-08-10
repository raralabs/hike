package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/raralabs/canal/ext/transforms/aggregates"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/transforms/agg"

	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/ext/transforms/doFn"
	csvsnk "github.com/raralabs/hike/plugins/canal/sinks"
	csvsrc "github.com/raralabs/hike/plugins/canal/sources"
	"github.com/raralabs/hike/utils/faker"
)

const TmpPath = "./tmp/"

func main() {
	if err := os.MkdirAll(TmpPath, os.ModePerm); err != nil {
		log.Panic(err)
	}

	p := pipeline.NewPipeline(1)
	opts := pipeline.DefaultProcessorOptions

	filename := TmpPath + "users.csv"
	//generateCsv(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}

	src := p.AddSource("CSV Reader")
	sp := src.AddProcessor(opts, csvsrc.NewCsvReader(file, -1))

	delay := p.AddTransform("Delay")
	del := delay.AddProcessor(opts, doFn.DelayFunction(10*time.Millisecond), "path1")

	//Count Last Names
	count := aggregates.NewCount("Count", func(m map[string]interface{}) bool {
		return true
	})
	after := func(m message.Msg, proc pipeline.IProcessorForExecutor, content, pContent []*message.OrderedContent) {

		contents := m.Content()
		if v, ok := contents.Get("eof"); ok {
			if v.Val == true {
				proc.Result(m, contents, nil)
				proc.Done()
				return
			}
		}
		for i := range content {
			proc.Result(m, content[i], pContent[i])
		}
	}
	aggs := []agg.IAggFuncTemplate{count}
	aggregator := agg.NewAggregator(aggs, after, "last_name")

	counter := p.AddTransform("Aggregator")
	cnt := counter.AddProcessor(opts, aggregator.Function(), "path2")

	//Count
	genAgg := aggregates.NewAvg("Generic", "Count", func(m map[string]interface{}) bool {
		return true
	})
	aggs2 := []agg.IAggFuncTemplate{genAgg}
	aggregator2 := agg.NewAggregator(aggs2, after)

	genericAgg := p.AddTransform("Generic Agg")
	gAgg := genericAgg.AddProcessor(opts, aggregator2.Function(), "path3")

	snk := p.AddSink("CSV Writer")
	snk.AddProcessor(opts, csvsnk.NewPrettyPrinter(os.Stdout, 100), "sink")

	delay.ReceiveFrom("path1", sp)
	counter.ReceiveFrom("path2", del)
	genericAgg.ReceiveFrom("path3", cnt)
	snk.ReceiveFrom("sink", gAgg)

	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	p.Validate()

	p.Start(c, cancel)

	file.Close()
}

func generateCsv(filename string) {

	file, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	// Generate Stream CSV Data
	choices := map[string][]interface{}{
		"first_name": {"Madhav", "Shambhu", "Pushpa", "Kumar", "Hero"},
		"last_name":  {"Mashima", "Dahal", "Poudel", "Rimal", "Amatya", "Shrestha", "Bajracharya"},
		"age":        {10, 20, 30, 40, 50, 60, 70, 15, 25, 35, 45, 55, 65, 75, 100, 6, 33, 47},
	}
	faker.FakeCsvData(file, choices, 10)
	file.Close()
}
