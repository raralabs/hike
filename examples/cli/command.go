package main

import (
	"context"
	"github.com/c-bata/go-prompt"
	"github.com/raralabs/hike/parser"
	"time"
)

func CommandMode(builder parser.IBuilder, history []string) []string {
	for {
		cmd := prompt.Input("Cmd>> ", cmdCompleter, prompt.OptionHistory(history))

		// Add command to history if it is a not the last command
		if cmd != history[len(history)-1] {
			if len(history) >= MaxHistoryLength {
				history = history[1:len(history)]
			}
			history = append(history, cmd)
		}

		if cmd == "end" {
			break
		}

		// Build the command
		starter, ppln, closer := builder.Build(1, cmd)

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
