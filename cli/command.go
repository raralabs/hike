package cli

import (
	"context"
	"fmt"
	"github.com/raralabs/hike/builder"
	"github.com/raralabs/hike/parser"
	"strings"
	"time"
)

func CommandMode(cmdFunc func(string) string, addHistory func(string), prsr parser.IParser) {
	commandBuilder := builder.NewCommand()
	pipelineId := uint32(1)

	promptText := "Cmd>> "
	var lastCommand string
	for {
		cmd := cmdFunc(promptText)
		str := strings.TrimSpace(cmd)
		fmt.Println(str)

		if str == "end" {
			break
		}

		if len(str) == 1 && str[0] == ';' {
			promptText = "Cmd>> "
			commandBuilder.Reset()
			addHistory(lastCommand + ";")
			continue
		}

		command, done, err := commandBuilder.Add(cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if done {
			cmd = command + ";"
			promptText = "Cmd>> "
		} else {
			cmd = command
			promptText = ">>>>> "
		}

		lastCommand = command
		addHistory(cmd)

		// Build the command
		starter, ppln, closer := prsr.Build(pipelineId, command)

		// Create context to run the pipeline
		c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		// Run the starter to initialize all stages of a pipeline
		starter()

		start := time.Now()

		// Run the pipeline
		ppln.Start(c, cancel)

		fmt.Printf("[TIME] To run the command: %v\n", time.Since(start))

		// Close everything
		closer()
	}
}
