package newMulti

import (
	"fmt"
	canalSrc "github.com/raralabs/canal/ext/sources"
	"os"

	"github.com/raralabs/canal/core/pipeline"

	"github.com/raralabs/canal/utils/cast"
	"github.com/raralabs/hike/parser/at"

	"github.com/raralabs/hike/plugins/canal/sources"
	"github.com/raralabs/hike/parser/newPeg"

	"log"


"sync"

)

type atBuilder struct {
	streamFromMu, streamToMu *sync.Mutex
	streamFrom               map[string][]at.Node
	streamTo                 map[string]*node
}

func NewATBuilder() *atBuilder {
	return &atBuilder{
		streamFromMu: &sync.Mutex{},
		streamToMu:   &sync.Mutex{},
		streamFrom:   make(map[string][]at.Node),
		streamTo:     make(map[string]*node),
	}
}

func (p *atBuilder) BuildAT(cmds []interface{}) at.AT{
	startId := int64(1)
	var srcs []at.Node
	for _,command := range cmds{
		var src at.Node
		src, startId = p.buildSinglePipeline(startId,cast.ToIfaceSlice(command))
		if src != nil{
			srcs = append(srcs,src)
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


func (p *atBuilder) buildSinglePipeline(Id int64,command []interface{})(at.Node, int64) {
	var firstNode *node
	for _, eachStage := range command {
		exec := getExecutor(eachStage)
		fmt.Println(firstNode,exec)
	}

		return firstNode, Id


}

func getNode(stg interface{},Id *int64 ) (*node){
	exec := getExecutor(stg)
	if exec == nil {
		log.Panic(stg)
	}
	n := &node{
		id:    *Id,
		exec:  exec,
		added: true,
	}
	*Id++
	return n
}

//generates all the source execcutors based on the types
//input is the stg of type SourceJob and output is the
//respective executor.
func getSourceExecutor(stg newPeg.SourceJob)pipeline.Executor{
	var exec pipeline.Executor
	switch stg.Type {
	case newPeg.FILEJOB:

		params := stg.OperateOn.(newPeg.SrcFile)
		fileName := params.FileName
		f, err := os.Open(fileName)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("[INFO] %s file opened as source.", fileName)
		exec = sources.NewCsvReader(f, -1)
		exec.SetName("file")

	case newPeg.FAKEJOB:
		params := stg.OperateOn.(newPeg.SrcFake)
		fmt.Println("iam here")
		numData := params.Number
		log.Printf("[INFO] Generating %d fake data", numData)
		exec = canalSrc.NewFaker(numData, nil)
		exec.SetName("fake")
	}
	return exec
}

func getExecutor(stg interface{}) pipeline.Executor {

	var exec pipeline.Executor
	switch stg.(type) {
	case newPeg.SourceJob:
		currentStg := stg.(newPeg.SourceJob)
		exec = getSourceExecutor(currentStg)

	case newPeg.TransformJob:
		fmt.Println("into the transform job")

	}
	return exec
}
