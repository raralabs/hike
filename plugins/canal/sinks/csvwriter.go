package sinks

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
)

type CsvWriter struct {
	name        string
	writer      *csv.Writer
	firstRecord bool
	header      []string
}

func NewCsvWriter(path string) *CsvWriter {

	// Create a file for writing
	f, err := os.Create(path)
	if err != nil {
		log.Panicf("Couldn't open file for writing. Error: %v", err)
	}

	return &CsvWriter{name: "FileWriter", writer: csv.NewWriter(f), firstRecord: true}
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
		record[i] = fmt.Sprintf("%v", content[cw.header[i]].Val)
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
