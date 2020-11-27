package newMulti

import (
	"fmt"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/parser/at"
	"strconv"
)

var (
	opts = pipeline.DefaultProcessorOptions
)



type netBuilder struct {
	nodeCounter map[at.Node]int32//keeps record of the nodes and its occurrence in the command
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
	 _,ok := nb.nodeCounter[node]

	 if ok {
		 if nb.nodeCounter[node] != 1 {

			 nb.nodeCounter[node] -= 1
			 nb.processorTable[node] = append(nb.processorTable[node], proc)
		 } else {
		 	 nb.nodeCounter[node] = 0
		 	 if exec.ExecutorType() == pipeline.SINK{
		 	 	 snk := p.AddSink(exec.Name())
		 	 	 snk.AddProcessor(opts,exec,"path1","path2")
				 nb.processorTable[node] = append(nb.processorTable[node],proc)
				 for idx, pcr := range nb.processorTable[node]{
					 route := "path" + strconv.Itoa(idx+1)
					 fmt.Println("get",route)
					 snk.ReceiveFrom(pipeline.MsgRouteParam(route),pcr)
				 }
				 return

			 }else if exec.ExecutorType() == pipeline.TRANSFORM{
				 transform := p.AddTransform(exec.Name())
				 pr := transform.AddProcessor(opts, exec, "path1", "path2")
				 nb.processorTable[node] = append(nb.processorTable[node],proc)
				 for idx, pcr := range nb.processorTable[node]{
					 route := "path" + strconv.Itoa(idx+1)
					 transform.ReceiveFrom(pipeline.MsgRouteParam(route),pcr)
				 }
				 for _, n := range node.ToNodes() {
					 nb.buildSubTree(n, p, pr)
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
			 pr := transform.AddProcessor(opts, exec, "path")
			 transform.ReceiveFrom("path", proc)
			 for _, n := range node.ToNodes() {
				 nb.buildSubTree(n, p, pr)
			 }

	 }
}