package sources

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/utils"
)

type CsvReader struct {
	name    string
	reader  *csv.Reader
	maxRows int
	currRow int
	header  []string
}

func NewCsvReader(r io.Reader, maxRows int) *CsvReader {

	csr := csv.NewReader(r)
	csr.TrimLeadingSpace = true

	header, err := csr.Read()
	if err != nil { //read header
		log.Panic(err)
	}

	return &CsvReader{
		reader:  csr,
		maxRows: maxRows,
		currRow: 0,
		header:  header,
	}
}

func (cr *CsvReader) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	if cr.currRow == cr.maxRows {
		cr.done(m, proc)
		return false
	}

	record, err := cr.reader.Read()
	if err == io.EOF {
		cr.done(m, proc)
		return false
	} else if err != nil {
		log.Panic(err)
	}

	if len(record) != len(cr.header) {
		log.Panic("Record and Header Length Mismatch")
	}

	content := message.NewOrderedContent()
	for i, v := range cr.header {
		value, valType := utils.GetValType(record[i])
		content.Add(v, message.NewFieldValue(value, valType))
	}
	proc.Result(m, content, nil)
	cr.currRow++

	return false
}

func (cr *CsvReader) done(m message.Msg, proc pipeline.IProcessorForExecutor) {
	// Send eof if done
	content := message.NewOrderedContent()
	content.Add("eof", message.NewFieldValue(true, message.BOOL))
	proc.Result(m, content, nil)
	proc.Done()
}

func (cr *CsvReader) ExecutorType() pipeline.ExecutorType {
	return pipeline.SOURCE
}

func (cr *CsvReader) HasLocalState() bool {
	return false
}

func (cr *CsvReader) SetName(name string) {
	cr.name = name
}

func (cr *CsvReader) Name() string {
	return cr.name
}
