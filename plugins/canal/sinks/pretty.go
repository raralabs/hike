package sinks

import (
	"fmt"
	"io"
	"strings"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"

	"github.com/jedib0t/go-pretty/table"
)

func stringRep(strs ...interface{}) string {
	var str strings.Builder

	for _, s := range strs {
		str.WriteString(fmt.Sprintf("%v", s))
	}

	return str.String()
}

type PrettyPrint struct {
	name        string
	writer      io.Writer
	firstRecord bool
	tw          table.Writer
	maxRows     uint64

	header  []string
	records map[string][]*message.MsgFieldValue
}

func NewPrettyPrinter(w io.Writer, maxRows uint64) *PrettyPrint {
	return &PrettyPrint{
		name:        "Pretty Printer",
		writer:      w,
		firstRecord: true,
		tw:          table.NewWriter(),
		maxRows:     maxRows,

		records: make(map[string][]*message.MsgFieldValue),
	}
}

func (cw *PrettyPrint) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	content := m.Content()

	if eof, ok := content.Get("eof"); ok {
		if eof.Val == true {

			row := make([]interface{}, len(cw.header))
			for i, h := range cw.header {
				row[i] = h
			}
			cw.tw.AppendHeader(row)

			numRecords := uint64(0)

			if len(cw.header) != 0 {
				depth := len(cw.records[cw.header[0]])

				for i := 0; i < depth; i++ {
					row := make([]interface{}, len(cw.records))
					j := 0
					for _, h := range cw.header {
						row[j] = cw.records[h][i].Value()
						j++
					}
					cw.tw.AppendRow(row)

					numRecords++
					if numRecords > cw.maxRows {
						break
					}
				}
			}

			cw.writer.Write([]byte(cw.tw.Render()))
			cw.writer.Write([]byte("\n"))

			if numRecords > cw.maxRows {
				cw.writer.Write([]byte("...and more.\n"))
			}

			proc.Done()
			return false
		}
	}

	if cw.firstRecord {
		cw.header = make([]string, content.Len())
		i := 0
		for e := content.First(); e != nil; e = e.Next() {
			k, _ := e.Value.(string)
			cw.header[i] = k
			i++
		}
		cw.firstRecord = false
	}

	replaced := false
	depth := 0
	if len(cw.header) != 0 {
		depth = len(cw.records[cw.header[0]])
	}

	pContent := m.PrevContent()
	if pContent != nil {
		for i := 0; i < depth; i++ {
			replace := true
			if content.Len() > pContent.Len() {
				for k, v := range cw.records {
					c, ok1 := content.Get(k)
					_, ok2 := pContent.Get(k)

					if ok1 && !ok2 {
						replace = replace && (v[i].Value() == c.Value() && v[i].ValueType() == c.ValueType())
						if replace {
							replaced = true
						} else {
							break
						}
					}
				}
			} else {
				replaced = true
			}

			if replace {
				for e := pContent.First(); e != nil; e = e.Next() {
					k, _ := e.Value.(string)

					if v, ok := cw.records[k]; ok {
						newVal, _ := content.Get(k)
						v[i] = newVal
					}
				}

			}
		}
	}
	if !replaced {
		for _, k := range cw.header {
			if val, ok := content.Get(k); ok {
				cw.records[k] = append(cw.records[k], val)
			}
		}
	}

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
