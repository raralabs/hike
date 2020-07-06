package parser

import "github.com/raralabs/canal/core/pipeline"

type IBuilder interface {
	// Init initializes the grammar and builds the parser.
	Init()

	// Build builds a pipeline on the basis of the provided command
	Build(id uint32, cmd string) (startFunc func(), ppln *pipeline.Pipeline, closeFunc func())
}
