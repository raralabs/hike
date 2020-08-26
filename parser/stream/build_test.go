package stream

import (
	"context"
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/ext/sinks"
	"github.com/raralabs/canal/ext/sources"
	"github.com/raralabs/canal/ext/transforms/doFn"
	"github.com/raralabs/hike/parser/at"
	"testing"
	"time"
)

type node struct {
	id int64
	exec pipeline.Executor
	toNodes []at.Node

	added bool
}

func (n *node) Add() {
	n.added = true
}

func (n *node) IsAdded() bool {
	return n.added
}

func (n *node) ID() int64 {
	return n.id
}

func (n *node) Executor() pipeline.Executor {
	return n.exec
}

func (n *node) ToNodes() []at.Node {
	return n.toNodes
}

type tree struct {
	sources []at.Node
}

func (t *tree) Sources() []at.Node {
	return t.sources
}

func TestBuild(t *testing.T) {

	src := &node{
		id:      1,
		exec:    sources.NewInlineRange(100),
	}

	proc1 := &node{
		id:      2,
		exec:    doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
			val := m["value"]
			if v, ok := val.(uint64); ok {
				if (v > 50) && (v%10 == 0) {
					return true, nil
				}
			}
			return false, nil
		}, func(msg message.Msg) bool {
			return false
		}),
	}

	proc2 := &node{
		id:      4,
		exec:    doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
			val := m["value"]
			if v, ok := val.(uint64); ok {
				if (v > 50) && (v%20 == 0) {
					return true, nil
				}
			}
			return false, nil
		}, func(msg message.Msg) bool {
			return false
		}),
	}

	snk1 := &node{
		id:      3,
		exec:    sinks.NewStdoutSink(),
	}

	snk2 := &node{
		id:      5,
		exec:    sinks.NewStdoutSink(),
	}

	src.toNodes = []at.Node{proc1, proc2}
	proc1.toNodes = []at.Node{snk1}
	proc2.toNodes = []at.Node{snk2}

	tr := &tree{sources: []at.Node{src}}

	p := Build(tr)
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	p.Start(c, cancel)
}
