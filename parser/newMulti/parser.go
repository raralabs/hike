package newMulti

import (
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/utils/cast"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
	"log"
	"strings"

	//"github.com/raralabs/hike/utils/LanguageProcessor"
	//"log"
)

//this is the structure for paser that is responsible for parsing the command
//as per the new peg rules and idioms
type parser struct {
	userCmd string
	stages []interface{}
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Init() {

	}

func IsComment(statement string)bool{
	if statement[0]=='#'{
		return true
	}
	return false
}

func (p *parser) Build(id uint32, cmds... string) (func(), *pipeline.Pipeline, func()) {
	var userCmd string
	userCmd = strings.Join(cmds, "\n")
	qryPreprocessor := LanguageProcessor.InitPreProcessor(userCmd)
	p.userCmd= qryPreprocessor.PrepareQuery(IsComment)
	parsedStages,err := newPeg.Parse("", []byte(p.userCmd))
	if err!= nil{
		log.Panic(err)
	}
	p.stages = cast.ToIfaceSlice(parsedStages)
	ab := NewATBuilder()
	nb := NewNetBuilder()
	tr := ab.Build(p.stages)
	ppln := nb.Build(id, tr)

	starter := func() {}
	closer := func() {}

	return starter, ppln, closer
}
