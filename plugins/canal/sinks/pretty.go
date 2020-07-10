package sinks

import (
	"io"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"

	"github.com/jedib0t/go-pretty/table"
)

type PrettyPrint struct {
	name        string
	writer      io.Writer
	firstRecord bool
	header      []string
	tw          table.Writer
	maxRows     uint64
	endMessage  bool
}

func NewPrettyPrinter(w io.Writer, maxRows uint64) *PrettyPrint {
	return &PrettyPrint{
		name:        "Pretty Printer",
		writer:      w,
		firstRecord: true,
		tw:          table.NewWriter(),
		maxRows:     maxRows,
		endMessage:  false,
	}
}

func (cw *PrettyPrint) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	content := m.Content()

	// Check for eof
	if v, ok := content.Get("eof"); ok {
		if v.Val == true {
			// Print the table
			cw.writer.Write([]byte(cw.tw.Render()))
			cw.writer.Write([]byte("\n"))
			if cw.endMessage {
				cw.writer.Write([]byte("...and more.\n"))
			}

			// All done
			proc.Done()
			return false
		}
	}

	if cw.maxRows == 0 {
		cw.endMessage = true
		return true
	}
	cw.maxRows--

	if cw.firstRecord {
		cw.header = make([]string, content.Len())
		row := make(table.Row, content.Len())
		i := 0
		for e := content.First(); e != nil; e = e.Next() {
			k, _ := e.Value.(string)
			cw.header[i] = k
			row[i] = k
			i++
		}
		cw.tw.AppendHeader(row)
		cw.firstRecord = false
	}

	row := make(table.Row, content.Len())
	for i := 0; i < len(cw.header); i++ {
		v, _ := content.Get(cw.header[i])
		if v == nil {
			row[i] = nil
		} else {
			row[i] = v.Val
		}
	}

	cw.tw.AppendRow(row)

	return false
}

func (cw *PrettyPrint) ExecutorType() pipeline.ExecutorType {
	return pipeline.SINK
}

func (cw *PrettyPrint) HasLocalState() bool {
	return false
}

func (cw *PrettyPrint) SetName(name string) {
	cw.name = name
}

func (cw *PrettyPrint) Name() string {
	return cw.name
}
