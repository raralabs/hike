package newMulti

import (
	"fmt"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/parser/at"
)

var (
	opts = pipeline.DefaultProcessorOptions
)

type JoinStat struct{
	JoinFlag bool
	firstProc pipeline.IProcessor
}

type netBuilder struct {
	JoinStatus JoinStat
}

func NewNetBuilder() *netBuilder {
	return &netBuilder{}
}

func (nb *netBuilder) Build(id uint32, tree at.AT) *pipeline.Pipeline {

	sources := tree.Sources()

	p := pipeline.NewPipeline(id)

	for _, s := range sources {
		src := p.AddSource(s.Executor().Name())
		sp := src.AddProcessor(opts, s.Executor())

		for _, n := range s.ToNodes() {
			fmt.Println("nodes",n.Executor().Name())
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

	//parse join node in a different way
	stageType := exec.Name()
	if  stageType == "join"{
		if nb.JoinStatus.JoinFlag==false{
			nb.JoinStatus.JoinFlag = true
			nb.JoinStatus.firstProc = proc

		}else{
			transform := p.AddTransform(exec.Name())
			pr:= transform.AddProcessor(opts, exec, "s1","s2")

			transform.ReceiveFrom("s1",nb.JoinStatus.firstProc)
			transform.ReceiveFrom("s2",proc)
			for _, n := range node.ToNodes(){
				nb.buildSubTree(n, p, pr)
			}
		}


	}else {
		// Parse the current node and build it's child
		transform := p.AddTransform(exec.Name())
		pr := transform.AddProcessor(opts, exec, "path")
		transform.ReceiveFrom("path", proc)
		for _, n := range node.ToNodes() {
			nb.buildSubTree(n, p, pr)

		}

	}
}