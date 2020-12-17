package newMulti

import (
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/parser/at"
	"strconv"
)

var (
	opts = pipeline.DefaultProcessorOptions
)

type netBuilder struct {
	nodeCounter map[at.Node]int32//keeps record of the number of subscription by any node.
	processorTable map[at.Node][]pipeline.IProcessor
}

func NewNetBuilder() *netBuilder {
	return &netBuilder{nodeCounter: make(map[at.Node]int32),processorTable: make(map[at.Node][]pipeline.IProcessor )}
}

func (nb *netBuilder) Build(id uint32, tree at.AT) *pipeline.Pipeline {

	sources := tree.Sources()
	nb.nodeCounter = tree.MultiNode()


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
	//parse join node in a different way
	// Parse the current node and build it's child
	 counter,ok := nb.nodeCounter[node]
	 if ok {
	 	if counter >1{
		 	 nb.nodeCounter[node] -= 1
			 nb.processorTable[node] = append(nb.processorTable[node], proc)
		 }else {
		 	 nb.nodeCounter[node] = 0
		 	 if exec.ExecutorType() == pipeline.SINK{
				 snk := p.AddSink(exec.Name())
		 	 	if nb.processorTable[node]!=nil{
		 	 	 snk.AddProcessor(opts,exec,"path1","path2")
				 nb.processorTable[node] = append(nb.processorTable[node],proc)
				 for idx, pcr := range nb.processorTable[node]{
					 route := "path" + strconv.Itoa(idx+1)
					 snk.ReceiveFrom(pipeline.MsgRouteParam(route),pcr)
				 }
		 	 	}else{
		 	 		snk.AddProcessor(opts,exec,"path1")
		 	 		snk.ReceiveFrom("path1",proc)
				}
				 return

			 }else if exec.ExecutorType() == pipeline.TRANSFORM{
				 transform := p.AddTransform(exec.Name())
				 if nb.processorTable[node]!=nil{
				 	pr := transform.AddProcessor(opts, exec, "path1", "path2")
				 	nb.processorTable[node] = append(nb.processorTable[node],proc)
				 	for idx, pcr := range nb.processorTable[node]{
					 	route := "path" + strconv.Itoa(idx+1)
					 	transform.ReceiveFrom(pipeline.MsgRouteParam(route),pcr)
				 	}
				 	for _, n := range node.ToNodes() {
					 	nb.buildSubTree(n, p, pr)
				 	}
				 }else{
					 pr:= transform.AddProcessor(opts,exec,"path1")
					 transform.ReceiveFrom("path1",proc)
					 for _, n := range node.ToNodes() {
						 nb.buildSubTree(n, p, pr)
					 }
				 }

			 }

			 //fmt.Println(nb.processorTable)
			 if exec.ExecutorType() == pipeline.SINK{
			 	return
			 }else{

			 }

		 }
		 }else{
		 	// Parse sink in a different way
		 	if exec.ExecutorType() == pipeline.SINK {
		 		snk := p.AddSink(exec.Name())
		 		snk.AddProcessor(opts, exec, "sink")
		 		snk.ReceiveFrom("sink", proc)
		 		return
		 	}

			 transform := p.AddTransform(exec.Name())

			 pr := transform.AddProcessor(opts, exec, "path1")
			 transform.ReceiveFrom("path1", proc)
			 for _, n := range node.ToNodes() {
				 nb.buildSubTree(n, p, pr)
			 }

	 }
}