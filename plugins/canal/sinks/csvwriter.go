package sinks

import (
	"encoding/csv"
	"fmt"
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"io"
	"strconv"
	"strings"
)

type CsvWriter struct {
	name        string
	writer      *csv.Writer
	firstRecord bool
	header      []string
}

func NewCsvWriter(w io.Writer, header ...string) *CsvWriter {
	return &CsvWriter{
		name:        "FileWriter",
		writer:      csv.NewWriter(w),
		firstRecord: true,
		header:      header,
	}
}

func (cw *CsvWriter) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	contents := m.Content()

	// Check for eof
	if v, ok := contents.Get("eof"); ok {
		if v.Val == true {
			proc.Done()
			return false
		}
	}

	if cw.firstRecord {
		if len(cw.header) == 0 {
			cw.header = contents.Keys()
		}
		cw.writer.Write(cw.header)
		cw.firstRecord = false
	}

	record := make([]string, contents.Len())
	for i := 0; i < len(cw.header); i++ {
		v, _ := contents.Get(cw.header[i])
		if v.Val == nil {
			record[i] = "nil"
		} else {
			if val, ok := v.Val.(float64); ok {
				s := strings.TrimRight(strconv.FormatFloat(val, 'f', 3, 64), "0")
				record[i] = strings.TrimRight(s, ".")
			} else {
				record[i] = fmt.Sprintf("%v", v.Val)
			}
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
