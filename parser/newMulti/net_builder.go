package newMulti

import (
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/parser/at"
)

var (
	opts = pipeline.DefaultProcessorOptions
)

type netBuilder struct {

}

func newNetBuilder() *netBuilder {
	return &netBuilder{}
}

func (nb *netBuilder) Build(id uint32, tree at.AT) *pipeline.Pipeline {

	sources := tree.Sources()

	p := pipeline.NewPipeline(id)

	for _, s := range sources {
		src := p.AddSource(s.Executor().Name())
		sp := src.AddProcessor(opts, s.Executor())

		for _, n := range s.ToNodes() {
			nb.buildSubTree(n, p, sp)
		}
	}

	p.Validate()

	return p
}

func (nb *netBuilder) buildSubTree(node at.Node, p *pipeline.Pipeline, proc pipeline.IProcessor) {

	exec := node.Executor()
	// Parse sink in a different way
	if exec.ExecutorType() == pipeline.SINK {
		snk := p.AddSink(exec.Name())
		snk.AddProcessor(opts, exec, "sink")
		snk.ReceiveFrom("sink", proc)
		return
	}

	// Parse the current node and build it's child
	transform := p.AddTransform(exec.Name())
	pr := transform.AddProcessor(opts, exec, "path")
	transform.ReceiveFrom("path", proc)

	for _, n := range node.ToNodes() {
		nb.buildSubTree(n, p, pr)
	}
}