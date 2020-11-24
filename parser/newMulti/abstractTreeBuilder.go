package newMulti

import (
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/utils/cast"
	"github.com/raralabs/hike/parser/at"
	"github.com/raralabs/hike/parser/newPeg"
	"log"
	"sync"
)

type executorPod struct{
	Executor 		pipeline.Executor
	PrimaryStage    bool
	StreamLabel 	interface{}
	StageType		string
}

//struct for the abstract tree builder
type atBuilder struct {
	streamFromMu, streamToMu *sync.Mutex// locks to avoid the race conditions
	streamFrom               map[string][]at.Node //keeps records of where the SECSOURCE will pass the msg
	streamTo                 map[string]*node //keeps records of where INTO sinks will pass the msg

}

//function that initializes the abstract tree builder
func NewATBuilder() *atBuilder {
	return &atBuilder{
		streamFromMu: &sync.Mutex{},
		streamToMu:   &sync.Mutex{},
		streamFrom:   make(map[string][]at.Node),
		streamTo:     make(map[string]*node),
	}
}

//takes the array of statements of commands parsed by the peg and constructs the abstract tree.
//abstract tree contains the node . the node contains toNodes which keeps track of the successive
//nodes to which it provides message to. (Similar to linked list)
func (p *atBuilder) BuildAT(cmds []interface{}) at.AT{
	startId := int64(1)
	var srcs []at.Node
	for _,statement := range cmds{
		var src at.Node
		src, startId = p.buildSinglePipeline(startId,cast.ToIfaceSlice(statement))
		if src != nil {
			srcs = append(srcs, src)
		}
	}

	p.streamToMu.Lock()
	p.streamFromMu.Lock()
	defer p.streamToMu.Unlock()
	defer p.streamFromMu.Unlock()

	for s := range p.streamTo {
		if tos, ok := p.streamFrom[s]; ok {
			p.streamTo[s].toNodes = append(p.streamTo[s].toNodes, tos...)
		}
	}

	absTree := &tree{sources: srcs}
	return absTree


}

//the function takes the startId and then starts to construct the executor and nodes for the
//single statement.
func (p *atBuilder) buildSinglePipeline(startId int64,statement []interface{})(at.Node, int64) {

	var firstNode *node //holds the first node. Mostly the source node of the statement
	var prevNode *node
	var multiStream []string
	streamFromName := ""
	streamToName := ""
	//traverse through the each stage of the statment src/transform*/sink and get
	//the respective executor if its a primary stage else get the other information
	//in the executor pod.
	for i,stage := range cast.ToIfaceSlice(statement){
		exec := getExecutor(stage)
		//file,fake,agg functions,stdout etc are primary stage
		//while into,branch etc are not.For primary stages executor
		//is needed. If executor is nil then panic.
		if exec.PrimaryStage == true{
			if exec.Executor == nil{
				log.Panic("Executor of primary stage is nil",stage)
			}
			//create a new node with the respective executor
			newNode := &node{
				id:    startId,
				exec:  exec.Executor,
				added: true,
			}
			//increment the id for next stage node
			startId++

			if exec.StageType == "SOURCE"{
				//if its a source and is at the first of the statement assign the node
				//to the firstNode else raise error
				if i==0{
					firstNode = newNode
				}else{
					log.Panic("source should always be the begining of the pipeline")
				}

			}else if exec.StageType == "TRANSFORM"|| exec.StageType == "SINK"{
				//if stage is transform and is the second stage of the statement and streamFromName exist
				//add it to the streamFrom.
				//streamFromName exist if the statement doesn't contain the
				//primary stage source.
				if i==1{
					if streamFromName != ""{
						p.streamFromMu.Lock()
						p.streamFrom[streamFromName] = append(p.streamFrom[streamFromName], newNode)
						p.streamFromMu.Unlock()
					} else if multiStream!=nil{
						p.streamFromMu.Lock()
						for _,secSource := range multiStream{
							p.streamFrom[secSource]= append(p.streamFrom[secSource],newNode)
						}
						p.streamFromMu.Unlock()
					}
				}

				if prevNode!=nil{
					prevNode.toNodes = append(prevNode.toNodes, newNode)
				}

			}
			prevNode = newNode
		}else if exec.PrimaryStage ==false{
			if exec.StageType == "SOURCE"{
				_,ok := exec.StreamLabel.(string)
					if ok{
						streamFromName = exec.StreamLabel.(string)
					}else{
						multiStream = exec.StreamLabel.([]string)
					}
				}
			}
			if exec.StageType == "SINK"{
				if prevNode != nil{
					streamToName = exec.StreamLabel.(string)
					p.streamToMu.Lock()
					if _, ok := p.streamTo[streamToName]; ok {
						p.streamToMu.Unlock()
						log.Panic("Can't have two different pipeline streaming to single stream")
					}
					p.streamTo[streamToName] = prevNode
					p.streamToMu.Unlock()
				}

			}
		}

	if firstNode==nil{
		return nil,startId
	}
	return firstNode, startId


}

//This function returns the executor based on the
//types of stage
func (p *atBuilder) getExecutor(stg interface{}) (pipeline.Executor) {

	var exec pipeline.Executor
	switch stg.(type) {
	//get the executor if stage is of source type
	case newPeg.SourceJob:
		currentStg := stg.(newPeg.SourceJob)
		exec = getSourceExecutor(currentStg)

	//get the executor if stage is of transform type
	case newPeg.TransformJob:
		currentStg := stg.(newPeg.TransformJob)
		exec = getTransformExecutor(currentStg)

	//get the executor if stage is of sink type
	case newPeg.SinkJob:
		currentStg := stg.(newPeg.SinkJob)
		exec = getSinkExecutor(currentStg)
	}
	return exec
}




