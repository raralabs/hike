package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/peterh/liner"
	"github.com/raralabs/hike/cli"
	"github.com/raralabs/hike/cli/cliner"
)

var ConsoleReader = bufio.NewReader(os.Stdin)

var (
	wsSuggestions = []string{
		"create",
		"open",
		"remove",
	}
)

func wsSuggest(line string) (c []string) {
	for _, n := range wsSuggestions {
		if strings.HasPrefix(n, strings.ToLower(line)) {
			c = append(c, n)
		}
	}
	return
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	line.SetCompleter(wsSuggest)

	cli.WSMode(func(prompt string) string {
		t, err := line.Prompt(prompt)
		if err != nil {
			log.Panic(err)
		}
		return t
	}, cliner.Liner, "newMulti")
}
