package sinks

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/raralabs/canal/core/transforms/agg"
	"github.com/raralabs/canal/utils/extract"

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

	batch *agg.Aggregator
	done  func(m message.Msg, proc pipeline.IProcessorForExecutor)
}

func NewPrettyPrinter(w io.Writer, maxRows uint64) *PrettyPrint {

	pp := &PrettyPrint{
		name:        "Pretty Printer",
		writer:      w,
		firstRecord: true,
		tw:          table.NewWriter(),
		maxRows:     maxRows,
	}

	return pp
}

func (cw *PrettyPrint) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

	//fmt.Println(m.Content(), m.PrevContent())

	if cw.firstRecord {
		content := m.Content()
		if content != nil {
			groups := extract.Columns(content)

			after := func(m message.Msg, proc pipeline.IProcessorForExecutor, contents, prevContents []*message.OrderedContent) {

				if content := m.Content(); content != nil {
					if eof, ok := content.Get("eof"); ok {
						if eof.Val == true {
							// Write header
							row := make([]interface{}, len(groups))
							for i, g := range groups {
								row[i] = g
							}
							cw.tw.AppendHeader(row)

							// Write maxRows entries
							numRecords := uint64(0)
							entries := cw.batch.Entries()
							for _, e := range entries {
								row := make([]interface{}, len(groups))
								for i, g := range groups {
									val, _ := e.Get(g)
									v := val.Value()

									var value interface{}
									if v == nil {
										value = "nil"
									} else if val, ok := v.(float64); ok {
										s := strings.TrimRight(strconv.FormatFloat(val, 'f', 3, 64), "0")
										value = strings.TrimRight(s, ".")
									} else {
										value = fmt.Sprintf("%v", v)
									}

									row[i] = value
								}
								cw.tw.AppendRow(row)

								numRecords++
								if numRecords > cw.maxRows {
									break
								}
							}

							cw.writer.Write([]byte(cw.tw.Render()))
							cw.writer.Write([]byte("\n"))

							if numRecords > cw.maxRows {
								cw.writer.Write([]byte("...and more.\n"))
							}

							proc.Done()
						}
					}
				}
			}

			cw.batch = agg.NewAggregator([]agg.IAggFuncTemplate{}, after, groups...)
			cw.firstRecord = false

			cw.done = func(m message.Msg, proc pipeline.IProcessorForExecutor) {
				after(m, proc, nil, nil)
			}
		}
	}

	if !cw.firstRecord {
		cw.batch.AggFunc(m, &struct{}{})
		cw.done(m, proc)
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
