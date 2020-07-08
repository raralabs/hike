package main

import (
	"context"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/raralabs/hike/builder"
	"github.com/raralabs/hike/parser"
	"strings"
	"time"
)

func CommandMode(prsr parser.IParser, history []string) []string {
	commandBuilder := builder.NewCommand()
	pipelineId := uint32(1)

	promptText := "Cmd>> "
	for {
		cmd := prompt.Input(promptText, cmdCompleter, prompt.OptionHistory(history))

		if cmd == "end" {
			break
		}

		str := strings.TrimSpace(cmd)
		if len(str) == 1 && str[0] == ';' {
			promptText = "Cmd>> "
			commandBuilder.Reset()
			lastCommand := history[len(history)-1]
			history = addHistory(history, lastCommand + ";")
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

		history = addHistory(history, cmd)

		// Build the command
		starter, ppln, closer := prsr.Build(pipelineId, command)

		// Create context to run the pipeline
		c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		// Run the starter to initialize all stages of a pipeline
		starter()

		// Run the pipeline
		ppln.Start(c, cancel)

		// Close everything
		closer()

	}

	return history
}

func addHistory(history []string, cmd string) []string {
	// Add command to history if it is a not the last command
	if cmd != history[len(history)-1] {
		if len(history) >= MaxHistoryLength {
			history = history[1:]
		}
		history = append(history, cmd)
	}

	return history
}
