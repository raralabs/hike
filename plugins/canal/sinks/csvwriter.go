package sinks

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
)

type CsvWriter struct {
	name        string
	writer      *csv.Writer
	firstRecord bool
	header      []string
}

func NewCsvWriter(w io.Writer) *CsvWriter {
	return &CsvWriter{name: "FileWriter", writer: csv.NewWriter(w), firstRecord: true}
}

func (cw *CsvWriter) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	content := m.Content()

	// Check for eof
	if v, ok := content["eof"]; ok {
		if v.Val == true {
			proc.Done()
			return false
		}
	}

	if cw.firstRecord {
		cw.header = make([]string, len(content))
		i := 0
		for k := range content {
			cw.header[i] = k
			i++
		}
		cw.writer.Write(cw.header)
		cw.firstRecord = false
	}

	record := make([]string, len(content))
	for i := 0; i < len(content); i++ {
		v := content[cw.header[i]]
		if v == nil {
			record[i] = "nil"
		} else {
			record[i] = fmt.Sprintf("%v", v.Val)
		}
	}

	cw.writer.Write(record)
	cw.writer.Flush()

	return false
}

func (cw *CsvWriter) ExecutorType() pipeline.ExecutorType {
	return pipeline.SINK
}

func (cw *CsvWriter) HasLocalState() bool {
	return false
}

func (cw *CsvWriter) SetName(name string) {
	cw.name = name
}

func (cw *CsvWriter) Name() string {
	return cw.name
}
