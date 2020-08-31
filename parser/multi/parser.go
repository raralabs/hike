package multi

import "github.com/raralabs/canal/core/pipeline"

type parser struct {
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Init() {
}

func (p *parser) Build(id uint32, cmds ...string) (func(), *pipeline.Pipeline, func()) {
	ab := newATBuilder()
	nb := newNetBuilder()

	tr := ab.Build(cmds...)
	ppln := nb.Build(id, tr)

	starter := func() {}
	closer := func() {}

	return starter, ppln, closer
}
