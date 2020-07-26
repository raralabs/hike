package main

import (
	"github.com/raralabs/go-wm/wm"
	"github.com/raralabs/hike/cli/cliner"
)

func main() {
	p := &wm.Project{
		History: ".tmp.cmd_history",
	}
	cliner.Liner(p)
}
