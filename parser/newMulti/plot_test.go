package newMulti

import (
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/ext/sinks"
	"github.com/raralabs/canal/ext/sources"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/hike/parser/at"
	"testing"
)

func TestPlot(t *testing.T) {

	src := &node{
		id:   1,
		exec: sources.NewInlineRange(100),
	}

	exec := doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
		val := m["value"]
		if v, ok := val.(uint64); ok {
			if (v > 50) && (v%10 == 0) {
				return true, nil
			}
		}
		return false, nil
	}, func(msg message.Msg) bool {
		return false
	})
	exec.SetName("Processor 1")

	proc1 := &node{
		id:   2,
		exec: exec,
	}

	exec = doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
		val := m["value"]
		if v, ok := val.(uint64); ok {
			if (v > 50) && (v%20 == 0) {
				return true, nil
			}
		}
		return false, nil
	}, func(msg message.Msg) bool {
		return false
	})
	exec.SetName("Processor 2")

	proc2 := &node{
		id:   4,
		exec: exec,
	}

	snk1 := &node{
		id:   3,
		exec: sinks.NewStdoutSink(),
	}

	snk2 := &node{
		id:   5,
		exec: sinks.NewStdoutSink(),
	}

	src.toNodes = []at.Node{proc1, proc2}
	proc1.toNodes = []at.Node{snk1}
	proc2.toNodes = []at.Node{snk2}

	tr := &tree{sources: []at.Node{src}}

	Plot(tr)
}
