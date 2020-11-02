package newMulti

import (
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/parser/at"
)

type node struct {
	id      int64
	exec    pipeline.Executor
	added   bool
	toNodes []at.Node
}

func (n *node) ID() int64 {
	return n.id
}

func (n *node) Executor() pipeline.Executor {
	return n.exec
}

func (n *node) Add() {
	n.added = true
}

func (n *node) IsAdded() bool {
	return n.added
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
