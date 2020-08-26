// Package `at` defines the Abstract Trees that is used by the
// parser to build the network.
package at

import "github.com/raralabs/canal/core/pipeline"

// Node represents a single stage in the pipeline
type Node interface {
	// ID must be unique for all nodes in an abstract tree.
	ID() int64

	// Executor returns the executor held by the Node.
	Executor() pipeline.Executor

	// Add adds a Node to the pipeline.
	Add()

	// IsAdded checks if the node has already been added to the network.
	IsAdded() bool

	// ToNodes returns all the nodes the current Node provides data to.
	ToNodes() []Node
}

// AT (or Abstract Tree) refers to a data structure consisting of
// Nodes which are loosely coupled.
// AT can be traversed recursively,
type AT interface {
	Sources() []Node
}
