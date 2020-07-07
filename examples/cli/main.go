package main

import (
	"context"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/raralabs/canal/utils/cast"
	"github.com/raralabs/hike/parser/peg"
	"github.com/raralabs/hike/utils/faker"
	"log"
	"os"
	"time"
)

const MaxHistoryLength = 20

func basicCompleter(d prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "faker", Description: "Generate fake csv file mode"},
		{Text: "log", Description: "Starts log mode"},
		{Text: "cmd", Description: "Starts command mode"},
		{Text: "end", Description: "Ends current mode"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func cmdCompleter(d prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "file(", Description: "Add a file source or sink"},
		{Text: "filter(", Description: "Add filter transform"},
		{Text: "count(", Description: "Add count aggregator"},
		{Text: "max(", Description: "Add max aggregator"},
		{Text: "min(", Description: "Add min aggregator"},
		{Text: "stdout(", Description: "Add stdout sink"},
		{Text: "end", Description: "Ends current mode"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func noneCompleter(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func HandlePanic(f func()) (err interface{}) {
	err = nil
	defer func() {
		if r := recover(); r != nil {
			err = r
		}
	}()

	f()

	return
}

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

	builder := peg.NewPegCmd()
	// Initialize the builder
	builder.Init()

	hikePrompt := func() {
		for {

			t := prompt.Input("> ", basicCompleter)

			log.Println("Here!!")

			switch t {
			case "cmd":
				// Command Mode
				for {
					cmd := prompt.Input("Cmd>> ", cmdCompleter, prompt.OptionHistory(commandHistory))

					// Add command to history if it is a new command
					if cmd != commandHistory[len(commandHistory)-1] {
						if len(commandHistory) >= MaxHistoryLength {
							commandHistory = commandHistory[1:len(commandHistory)]
						}
						commandHistory = append(commandHistory, cmd)
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

			case "faker":
				// Fake Data Generation Mode
				fileName := prompt.Input("Filename>> ", noneCompleter)
				numRows := prompt.Input("NumRows>> ", noneCompleter)
				if rows, ok := cast.TryInt(numRows); ok {
					generateCsv(fileName, int(rows))
					fmt.Printf("Fake csv file generation success. File: %s\n", fileName)
				} else {
					fmt.Println("Fake csv file generation failed.")
				}

			case "log":
				// Log File Config Mode
				fileName := prompt.Input("Log Filename>> ", noneCompleter)
				closer := setLogFile(fileName)
				fmt.Printf("Log file config success. File: %s\n", fileName)

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
			log.Printf("Recovered: %v", err)
		}
	}
}

func setLogFile(filename string) func() {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	return func() {
		f.Close()
	}
}

func generateCsv(filename string, rows int) {

	file, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	// Generate Fake CSV Data
	choices := map[string][]interface{}{
		"first_name": {"Madhav", "Shambhu", "Pushpa", "Kumar", "Hero"},
		"last_name":  {"Mashima", "Dahal", "Poudel", "Rimal", "Amatya", "Shrestha", "Bajracharya"},
		"age":        {10, 20, 30, 40, 50, 60, 70, 15, 25, 35, 45, 55, 65, 75, 100, 6, 33, 47},
	}
	faker.FakeCsvData(file, choices, rows)
	file.Close()

	log.Println("CSV File Generation Complete")
	log.Printf("Total Rows: \t%v", rows)
}
