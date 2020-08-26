package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/raralabs/hike/builder"
	"github.com/raralabs/hike/parser"
)

func CommandMode(cmdFunc func(string) string, addHistory func(string), prsr parser.IParser) {
	commandBuilder := builder.NewCommand()
	streamBuilder := builder.NewStreamCommand()
	pipelineId := uint32(1)

	promptText := "Cmd>> "
	var lastCommand string

	var multiCommands []string
	for {
		cmd := cmdFunc(promptText)
		str := strings.TrimSpace(cmd)

		if str == "end" {
			break
		}

		if str == ";" {
			promptText = "Cmd>> "
			commandBuilder.Reset()
			addHistory(lastCommand + ";")
			continue
		} else if str == ";;" {
			printCommands(multiCommands...)
			multiCommands = nil

			promptText = "Cmd>> "
			commandBuilder.Reset()
			addHistory(lastCommand + ";")

			continue
		}

		stream, err := streamBuilder.Build(cmd)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		command, done, err := commandBuilder.Add(stream)
		if err != nil {
			fmt.Println(err)
			continue
		}

		addHistory(cmd)
		if done {
			multiCommands = append(multiCommands, cmd)
			cmd = command + ";"

			promptText = "Cmd>> "
		} else {
			multiCommands = append(multiCommands, cmd)
			cmd = command

			promptText = ">>>>> "
		}

		lastCommand = command

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

func printCommands(cmds ...string) {
	fmt.Println("Commands")
	for _, c := range cmds {
		fmt.Println(c)
	}
}