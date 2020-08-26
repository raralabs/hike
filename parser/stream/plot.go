package stream

import (
	"fmt"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/parser/at"
	"strings"
)

func Plot(tree at.AT) {

	sources := tree.Sources()
	depth := 1

	for _, s := range sources {
		fmt.Printf("%s", s.Executor().Name())

		for i, n := range s.ToNodes() {
			if i > 0 {
				for i := 0; i < depth; i++ {
					fmt.Print("  ")
				}
			}
			fmt.Println(plotSubTree(n, depth+1))
		}
	}
}

func plotSubTree(node at.Node, depth int) string {
	var str strings.Builder

	exec := node.Executor()
	// Parse sink in a different way
	if exec.ExecutorType() == pipeline.SINK {
		return " -> " + exec.Name()
	}

	str.WriteString(" -> ")
	str.WriteString(exec.Name())

	for i, n := range node.ToNodes() {
		if i > 0 {
			str.WriteString("\n")
			for i := 0; i < depth; i++ {
				str.WriteString("    ")
			}
		}
		str.WriteString(plotSubTree(n, depth+1))
	}

	return str.String()
}
