package parser

import "github.com/raralabs/canal/core/pipeline"

type IParser interface {
	// Init initializes the grammar and builds the parser.
	Init()

	// Build builds a pipeline on the basis of the provided command
	Build(id uint32, cmds ...string) (startFunc func(), ppln *pipeline.Pipeline, closeFunc func())
}
