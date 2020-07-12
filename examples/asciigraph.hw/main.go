package main

import (
	"fmt"

	"github.com/raralabs/hike/plugins/asciigraph/plot"
)

func main() {
	data := []float64{
		40, 20, 30,
	}

	bar := plot.NewBar()
	bar.Data = data
	bar.Labels = []string{"Aquarium", "Bacteria", "Camphor"}
	bar.Width = 2
	bar.Gap = 5

	graph := bar.Plot()

	fmt.Println(graph)
}
