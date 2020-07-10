package main

import (
	"fmt"

	"github.com/raralabs/hike/cli"

	"github.com/c-bata/go-prompt"
	"github.com/raralabs/hike/parser/peg"
)

const MaxHistoryLength = 20

func main() {
	fmt.Println("Please enter command to run. (Remember to set log file)")

	var closeFuncs []func()
	defer func() {
		for _, fn := range closeFuncs {
			fn()
		}
	}()

	if MaxHistoryLength < 2 {
		panic("Can't have history less than 2")
	}
	commandHistory := make([]string, MaxHistoryLength)
	prsr := peg.NewPegCmd()

	// Initialize the parser
	prsr.Init()

	hikePrompt := func() {
		for {

			t := prompt.Input("> ", BasicCompleter)

			switch t {
			case "cmd":
				// Command Mode
				cli.CommandMode(func(prmpt string) string {
					return prompt.Input(prmpt, cmdCompleter, prompt.OptionHistory(commandHistory))
				}, func(cmd string) {
					// Add command to history if it is a not the last command
					if cmd != commandHistory[len(commandHistory)-1] {
						if len(commandHistory) >= MaxHistoryLength {
							commandHistory = commandHistory[1:]
						}
						commandHistory = append(commandHistory, cmd)
					}
				}, prsr)

			case "faker":
				// Fake Data Generation Mode
				cli.FakerMode(func(prmpt string) string {
					return prompt.Input(prmpt, noneCompleter)
				})

			case "log":
				// Log File Config Mode
				closer := cli.LogMode(func(prmpt string) string {
					return prompt.Input(prmpt, noneCompleter)
				})
				closeFuncs = append(closeFuncs, closer)
			}

			if t == "end" {
				break
			}
		}
	}

	for {
		if err := cli.HandlePanic(hikePrompt); err == nil {
			break
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
