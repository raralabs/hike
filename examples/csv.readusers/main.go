package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/raralabs/canal/ext/transforms/aggregates"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/core/transforms/agg"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/canal/utils/cast"

	csvsnk "github.com/raralabs/hike/plugins/canal/sinks"
	csvsrc "github.com/raralabs/hike/plugins/canal/sources"
	"github.com/raralabs/hike/utils/faker"
)

const TmpPath = "./tmp/"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	ageFilt := func(m map[string]interface{}) (bool, error) {
		match := false

		if rawAge, ok := m["age"]; ok {
			if age, ok := cast.TryFloat(rawAge); ok {
				if age > 20 && age < 50 {
					match = true
				}
			}
		}

		return match, nil
	}

	filt := p.AddTransform("Age Filter")
	f1 := filt.AddProcessor(opts, doFn.FilterFunction(ageFilt, doneFunc), "path2")

	// Count Last Names
	count := aggregates.NewCount("Count", func(m map[string]interface{}) bool {
		return true
	})
	aggs := []agg.IAggFuncTemplate{count}
	aggregator := agg.NewAggregator(aggs, nil, "last_name")

	counter := p.AddTransform("Aggregator")
	cnt := counter.AddProcessor(opts, aggregator.Function(), "path3")

	snk := p.AddSink("CSV Writer")
	snk.AddProcessor(opts, csvsnk.NewCsvWriter(os.Stdout), "sink")

	delay.ReceiveFrom("path1", sp)
	filt.ReceiveFrom("path2", del)
	counter.ReceiveFrom("path3", f1)
	snk.ReceiveFrom("sink", cnt)

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
