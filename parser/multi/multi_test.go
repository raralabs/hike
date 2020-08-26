package multi

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
	"testing"
	"time"
)

//func TestBuild(t *testing.T) {
//	cmds := []string{
//		"fake(5) | select(age, last_name) into s1",
//		"s1 | stdout()",
//		"s1 | filter(age > 20) | stdout()",
//		"s1 | map twice_age = 2 * age | stdout()",
//		"s1 | map half_age = age / 2 | stdout()",
//	}
//	absTree := Build(cmds...)
//	stream.Plot(absTree)
//
//	p := stream.Build(absTree)
//	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
//
//	p.Start(c, cancel)
//}
//
//func TestEngine(t *testing.T) {
//
//	src := &node{
//		id:       1,
//		executor: sources.NewFaker(10, nil),
//	}
//
//	proc1 := &node{
//		id: 2,
//		executor: doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
//			return true, nil
//		}, func(msg message.Msg) bool {
//			return false
//		}),
//	}
//
//	expr, err := govaluate.NewEvaluableExpression("age * 2")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	proc2 := &node{
//		id: 4,
//		executor: doFn.EnrichFunction("twice_age", expr, func(msg message.Msg) bool {
//			return false
//		}),
//	}
//
//	snk1 := &node{
//		id:       3,
//		executor: sinks.NewPrettyPrinter(os.Stdout, 10),
//	}
//
//	snk2 := &node{
//		id:       5,
//		executor: sinks.NewPrettyPrinter(os.Stdout, 5),
//	}
//
//	src.toNodes = []at.Node{proc1, proc2}
//	proc1.toNodes = []at.Node{snk1}
//	proc2.toNodes = []at.Node{snk2}
//
//	tr := &tree{sources: []at.Node{src}}
//
//	p := stream.Build(tr)
//	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
//
//	p.Start(c, cancel)
//}

func TestEngineRaw(t *testing.T) {
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
