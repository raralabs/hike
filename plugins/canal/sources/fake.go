package sources

import (
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/utils"
	"github.com/raralabs/hike/utils/faker"
)

type Fake struct {
	name    string
	maxRows int64
	currRow int64
	faker   *faker.Stream
}

func NewFaker(maxRows int64, m map[string][]interface{}) *Fake {

	if m == nil {
		m = faker.DefaultChoices
	}

	return &Fake{
		name:    "Fake",
		maxRows: maxRows,
		currRow: 0,
		faker:   faker.NewFake(m),
	}
}

func (cr *Fake) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	if cr.currRow == cr.maxRows {
		cr.done(m, proc)
		return false
	}

	content := message.NewOrderedContent()
	fakeData := cr.faker.Random()

	for k, v := range fakeData {
		value, valType := utils.GetValType(v)
		content.Add(k, message.NewFieldValue(value, valType))
	}
	proc.Result(m, content, nil)
	cr.currRow++

	return false
}

func (cr *Fake) done(m message.Msg, proc pipeline.IProcessorForExecutor) {
	// Send eof if done
	content := message.NewOrderedContent()
	content.Add("eof", message.NewFieldValue(true, message.BOOL))
	proc.Result(m, content, nil)
	proc.Done()
}

func (cr *Fake) ExecutorType() pipeline.ExecutorType {
	return pipeline.SOURCE
}

func (cr *Fake) HasLocalState() bool {
	return false
}

func (cr *Fake) SetName(name string) {
	cr.name = name
}

func (cr *Fake) Name() string {
	return cr.name
}
