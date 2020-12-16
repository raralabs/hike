package sinks

import (
	"fmt"
	"github.com/raralabs/canal/core/message/content"
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
	header      []string

	batch *agg.Aggregator
	done  func(m message.Msg, proc pipeline.IProcessorForExecutor)
}

func NewPrettyPrinter(w io.Writer, maxRows uint64, header ...string) *PrettyPrint {

	pp := &PrettyPrint{
		name:        "Pretty Printer",
		writer:      w,
		firstRecord: true,
		tw:          table.NewWriter(),
		maxRows:     maxRows,
		header:      header,
	}

	return pp
}

func (cw *PrettyPrint) Execute(m pipeline.MsgPod, proc pipeline.IProcessorForExecutor) bool {

	//fmt.Println(m.Content(), m.PrevContent())
	if cw.firstRecord {
		cntnt := m.Msg.Content()

		if cntnt != nil {
			var groups []string
			if cw.header == nil || len(cw.header) == 0 {
				groups = extract.Columns(cntnt)
			} else {
				groups = cw.header
			}

			after := func(m message.Msg, proc pipeline.IProcessorForExecutor, contents, prevContents []content.IContent) {
				if cntnt := m.Content(); cntnt != nil {
					if eof, ok := cntnt.Get("eof"); ok {
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
		cw.batch.AggFunc(m.Msg, &struct{}{})
		cw.done(m.Msg, proc)
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
