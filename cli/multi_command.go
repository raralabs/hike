package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/raralabs/hike/parser"
)

func MultiCommandMode(cmdFunc func(string) string, addHistory func(string), prsr parser.IParser) {
	pipelineId := uint32(1)

	promptText := "Cmd>> "
	var commands []string

	for {
		cmd := cmdFunc(promptText)
		str := strings.TrimSpace(cmd)

		if str == "end" {
			break
		}

		str = appendSink(str)

		var done bool
		if str == ";" {
			promptText = "Cmd>> "

			if len(commands) == 0 {
				continue
			}
			done = true
		} else {
			promptText = ">>>>> "
			if str[len(str)-1] == ';' {
				promptText = "Cmd>> "
				done = true
			}
			addHistory(str)

			commands = append(commands, str)
		}

		if done {
			// Build the command
			starter, ppln, closer := prsr.Build(pipelineId, commands...)

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

			// Reset the commands
			commands = nil
		}
	}
}

var (
	expectedSinks = []string{
		"stdout", "blackhole",
	}
)

// appendSink examines the command and see if it needs to
// add a sink. An stdout() sink is added if it's required.
func appendSink(cmd string) string {

	// If it's just an ending command, return it
	if cmd == ";" {
		return cmd
	}

	var endless string
	var ended bool

	// Check if the command is last command.
	if cmd[len(cmd)-1] == ';' {
		endless = cmd[:len(cmd)-1]
		ended = true
	} else {
		// If not an end command, needs to take care of into
		fields := strings.Fields(cmd)
		if len(fields) >= 2 {
			if fields[len(fields)-2] == "into" {
				return cmd
			}
		}
		endless = cmd
	}
	endless = strings.TrimSpace(endless)

	// Add a sink if it's not there
	fields := strings.Fields(endless)
	sink := fields[len(fields)-1]

	snkName := strings.Split(sink, "(")[0]
	if strInSlice(expectedSinks, snkName) {
		return cmd
	}

	// Append an stdout() sink
	newCmd := fmt.Sprintf("%s | stdout()", endless)
	if ended {
		newCmd = fmt.Sprintf("%s;", newCmd)
	}
	return newCmd
}

func strInSlice(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func printCommands(cmds ...string) {
	fmt.Println("Commands")
	for _, c := range cmds {
		fmt.Println(c)
	}
}
