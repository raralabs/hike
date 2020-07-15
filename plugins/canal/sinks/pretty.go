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
	//fmt.Println(content, m.PrevContent())

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
						v := cw.records[h][i].Value()
						if v == nil {
							v = "nil"
						}
						row[j] = v
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

	depth := 0
	if len(cw.header) != 0 {
		depth = len(cw.records[cw.header[0]])
	}

	pContent := m.PrevContent()
	if pContent != nil {
		for i := 0; i < depth; i++ {
			replace := true
			for _, h := range cw.header {
				v1, ok1 := cw.records[h]
				v2, ok2 := pContent.Get(h)

				if !ok1 || !ok2 {
					goto breakLoop
				}

				if !(v1[i].Value() == v2.Value() && v1[i].ValueType() == v2.ValueType()) {
					replace = false
				}

				if !replace {
					break
				}
			}

			if replace {
				for _, h := range cw.header {
					val, _ := content.Get(h)
					cw.records[h][i] = val
				}
				return false
			}
		}
	}
breakLoop:
	for _, k := range cw.header {
		if val, ok := content.Get(k); ok {
			cw.records[k] = append(cw.records[k], val)
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
