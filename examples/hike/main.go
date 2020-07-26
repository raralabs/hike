package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/raralabs/hike/cli"
	"github.com/raralabs/hike/cli/cliner"
)

var ConsoleReader = bufio.NewReader(os.Stdin)

func main() {
	cli.WSMode(prompt, cliner.Liner)
}

func prompt(str string) string {
	fmt.Print(str)
	text, _ := ConsoleReader.ReadString('\n')
	return strings.TrimSpace(text[:len(text)-1])
}
