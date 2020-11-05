package main

import (
"github.com/raralabs/canal/core/pipeline"
"github.com/raralabs/hike/parser/at"
)

//represents the stages of the pipeline
type node struct {
	id      int64
	exec    pipeline.Executor
	added   bool
	toNodes []at.Node
}

//get the id of the node
func (n *node) ID() int64 {
	return n.id
}

//return the executor of the pipeline
func (n *node) Executor() pipeline.Executor {
	return n.exec
}

//commits the addition of executor in the node
func (n *node) Add() {
	n.added = true
}

//returns the status of the executor added or not added to the node
func (n *node) IsAdded() bool {
	return n.added
}

//returns the nodes where the executor should stream the data
func (n *node) ToNodes() []at.Node {
	return n.toNodes
}

type tree struct {
	sources []at.Node
}

func (t *tree) Sources() []at.Node {
	return t.sources
}

