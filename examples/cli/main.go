package main

import (
	"fmt"
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

			t := prompt.Input("> ", basicCompleter)

			switch t {
			case "cmd":
				// Command Mode
				commandHistory = CommandMode(prsr, commandHistory)

			case "faker":
				// Fake Data Generation Mode
				FakerMode()

			case "log":
				// Log File Config Mode
				closer := LogMode()
				closeFuncs = append(closeFuncs, closer)
			}

			if t == "end" {
				break
			}
		}
	}

	for {
		if err := HandlePanic(hikePrompt); err == nil {
			break
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
