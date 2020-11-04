package multi

import (
	"context"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/ext/sources"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/hike/parser/at"
	"github.com/raralabs/hike/plugins/canal/sinks"
	"log"
	"os"
	"testing"
	"time"
)

func TestATBuild(t *testing.T) {
	cmds := []string{
		"s1 | filter(age > 30) | filter(age>23) | filter(age<10) | stdout()",
		"s1 | map twice_age = 2 * age | stdout()",
		"s1 | map half_age = age / 2 | stdout()",
	}

	builder := newATBuilder()
	absTree := builder.Build(cmds...)
	fmt.Println("Tree",absTree)
	Plot(absTree)
	nb := newNetBuilder()

	p := nb.Build(1, absTree)
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	p.Start(c, cancel)
}

func TestEngine(t *testing.T) {

	src := &node{
		id:   1,
		exec: sources.NewFaker(10, nil),
	}

	proc1 := &node{
		id: 2,
		exec: doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
			return true, nil
		}, func(msg message.Msg) bool {
			return false
		}),
	}

	expr, err := govaluate.NewEvaluableExpression("age * 2")
	if err != nil {
		log.Println(err)
		return
	}
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
	proc2 := &node{
		id:   4,
		exec: doFn.EnrichFunction("twice_age", expr, doneFunc),
	}

	snk1 := &node{
		id:   3,
		exec: sinks.NewPrettyPrinter(os.Stdout, 10),
	}

	snk2 := &node{
		id:   5,
		exec: sinks.NewPrettyPrinter(os.Stdout, 5),
	}

	src.toNodes = []at.Node{proc1, proc2}
	proc1.toNodes = []at.Node{snk1}
	proc2.toNodes = []at.Node{snk2}

	tr := &tree{sources: []at.Node{src}}
	nb := newNetBuilder()

	p := nb.Build(1, tr)
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	p.Start(c, cancel)
}
