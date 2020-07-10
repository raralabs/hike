package sinks

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/canal/utils/cast"
)

type BarPlot struct {
	name   string
	xField string
	yField string
	bc     *widgets.BarChart
	data   []float64
	labels []string
}

func NewBarPlot(title, xField, yField string, width, gap int) *BarPlot {

	bc := widgets.NewBarChart()
	bc.Title = title
	bc.BarWidth = width
	bc.BarGap = gap
	return &BarPlot{
		name:   "Bar Plot",
		xField: xField,
		yField: yField,
		bc:     bc,
	}
}

func (cw *BarPlot) Execute(m message.Msg, proc pipeline.IProcessorForExecutor) bool {

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

			//Plot the data
			if err := ui.Init(); err != nil {
				log.Panicf("Failed to initialize termui: %v", err)
			}
			defer ui.Close()

			cw.bc.Data = cw.data
			cw.bc.Labels = cw.labels
			cw.bc.SetRect(0, 0, len(cw.bc.Data)*(cw.bc.BarWidth+cw.bc.BarGap), 10)

			ui.Render(cw.bc)

			uiEvents := ui.PollEvents()
			for {
				e := <-uiEvents
				switch e.ID {
				case "q", "<C-c>":
					// All done
					proc.Done()
					return false
				}
			}
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
