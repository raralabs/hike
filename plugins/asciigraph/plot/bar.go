package plot

import (
	"log"
	"math"
	"strings"

	"github.com/guptarohit/asciigraph"
)

type Bar struct {
	Width  int
	Gap    int
	Labels []string
	Data   []float64
	Height int
}

func NewBar() *Bar {
	return &Bar{
		Width:  2,
		Gap:    5,
		Labels: nil,
		Data:   nil,
		Height: 10,
	}
}

func (b *Bar) Plot() string {
	// Check if data are available
	noData := false
	if len(b.Labels) == 0 {
		noData = true
		log.Println("[ERROR] Labels not available")
	}
	if len(b.Data) == 0 {
		noData = true
		log.Println("[ERROR] Data not available")
	}
	if noData {
		return ""
	}

	data := b.adjustData()
	caption := b.adjustLabel()

	graph := asciigraph.Plot(data,
		asciigraph.Height(b.Height),
		asciigraph.Caption(caption),
	)

	return graph
}

func (b *Bar) adjustData() []float64 {
	data := make([]float64, b.Gap+1)

	for i := 0; i < b.Gap+1; i++ {
		data[i] = 0
	}

	for _, d := range b.Data {
		dt := make([]float64, b.Width+b.Gap+2)
		i := 0
		for ; i < b.Width+1; i++ {
			dt[i] = d
		}

		for ; i < b.Gap+1; i++ {
			dt[i] = 0
		}

		data = append(data, dt...)
	}

	return data
}

func (b *Bar) adjustLabel() string {
	sepString := "_"

	var caption strings.Builder
	spc := int(math.Floor(float64(b.Gap) / 2))

	for i := 0; i < spc-1; i++ {
		caption.WriteString(sepString)
	}

	frameSz := b.Width + b.Gap + 1

	for _, l := range b.Labels {
		caption.WriteString(frame(l, sepString, frameSz))
		caption.WriteString(sepString)
	}

	return caption.String()
}

func frame(str, sepString string, frameSz int) string {

	if len(str) < frameSz {
		var out strings.Builder
		diff := frameSz - len(str)
		spc := int(math.Floor(float64(diff) / 2))

		for i := 0; i < spc; i++ {
			out.WriteString(sepString)
		}
		out.WriteString(str)

		if diff%2 == 0 {
			for i := 0; i < spc; i++ {
				out.WriteString(sepString)
			}
		} else {
			for i := 0; i < spc+1; i++ {
				out.WriteString(sepString)
			}
		}
		return out.String()
	}

	return str[:frameSz]
}
