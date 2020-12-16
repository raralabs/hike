package sinks

import (
	"fmt"
	"log"

	"github.com/raralabs/hike/plugins/asciigraph/plot"

	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/utils/cast"
)

type BarPlot struct {
	name   string
	xField string
	yField string
	bc     *plot.Bar
	data   []float64
	labels []string
}

func NewBarPlot(title, xField, yField string, width, gap int) *BarPlot {

	bc := plot.NewBar()
	bc.Width = width
	bc.Gap = gap
	return &BarPlot{
		name:   "Bar Plot",
		xField: xField,
		yField: yField,
		bc:     bc,
	}
}

func (cw *BarPlot) Execute(msgPod pipeline.MsgPod, proc pipeline.IProcessorForExecutor) bool {
	m := msgPod.Msg
	content := m.Content()

	// Check for eof
	if v, ok := content.Get("eof"); ok {
		if v.Val == true {
			// Simply return if no data is available
			noData := false
			if len(cw.data) == 0 {
				log.Println("[WARN] No y-Axis data to plot")
				noData = true
			}
			if len(cw.labels) == 0 {
				log.Println("[WARN] No x-Axis data to plot")
				noData = true
			}
			if noData {
				proc.Done()
				return false
			}

			cw.bc.Data = cw.data
			cw.bc.Labels = cw.labels

			graph := cw.bc.Plot()
			fmt.Println(graph)

			proc.Done()
			return false
		}
	}

	if x, ok := content.Get(cw.xField); ok {
		xs := fmt.Sprintf("%v", x.Value())
		cw.labels = append(cw.labels, xs)
	}

	if y, ok := content.Get(cw.yField); ok {
		if yf, ok := cast.TryFloat(y.Value()); ok {
			cw.data = append(cw.data, yf)
		}
	}

	return false
}

func (cw *BarPlot) ExecutorType() pipeline.ExecutorType {
	return pipeline.SINK
}

func (cw *BarPlot) HasLocalState() bool {
	return false
}

func (cw *BarPlot) SetName(name string) {
	cw.name = name
}

func (cw *BarPlot) Name() string {
	return cw.name
}
