package cliner

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/peterh/liner"
	"github.com/raralabs/go-wm/wm"
	"github.com/raralabs/hike/cli"
	"github.com/raralabs/hike/parser/peg"
)

var (
	defaultSuggestions = []string{
		"cmd",
		"faker",
		"log",
		"end",
	}

	commandSuggestions = []string{
		"file(",

		"filter(",
		"select(",

		"pick ",
		"first(", "last(", "random(",

		"count(",
		"dcount(",
		"max(",
		"min(",
		"avg(",
		"sum(",
		"var(",
		"median(",
		"quartile(",
		"quantile(",

		"plot ",
		"bar(",

		"end",
	}
)

func noneSuggest(line string) (c []string) {
	return
}

func defaultSuggest(line string) (c []string) {
	for _, n := range defaultSuggestions {
		if strings.HasPrefix(n, strings.ToLower(line)) {
			c = append(c, n)
		}
	}
	return
}

func commandSuggest(line string) (c []string) {
	words := strings.Fields(line)
	lastWord := words[len(words)-1]
	for _, n := range commandSuggestions {
		if strings.HasPrefix(n, strings.ToLower(lastWord)) {
			curr := strings.Join(words[:len(words)-1], " ")
			suggest := fmt.Sprintf("%s %s", curr, n)
			c = append(c, suggest)
		}
	}
	return
}

func Liner(p *wm.Project) {
	line := liner.NewLiner()
	defer line.Close()

	var closeFuncs []func()
	defer func() {
		for _, fn := range closeFuncs {
			fn()
		}

		if f, err := os.Create(p.History); err != nil {
			log.Print("Error writing history file: ", err)
		} else {
			line.WriteHistory(f)
			f.Close()
		}
	}()

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	if f, err := os.Open(p.History); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	prsr := peg.NewPegCmd()

	// Initialize the parser
	prsr.Init()

	// Set the log file
	t, err := line.Prompt("Set Log File?(y/n) ")
	if err == liner.ErrPromptAborted {
		fmt.Println("Aborted")
		return
	} else if err != nil {
		log.Panic(err)
	}

	if strings.ToLower(t) == "y" {
		closer := cli.LogMode(func(prompt string) string {
			t, err := line.Prompt(prompt)
			if err != nil {
				log.Panic(err)
			}
			return t
		})
		closeFuncs = append(closeFuncs, closer)
	} else {
		cli.SetLogFile("default.tmp.log")
		fmt.Println("Default logging to: default.tmp.log")
	}

	hikePrompt := func() {
		defaultSuggestionsSet := false
		for {
			if !defaultSuggestionsSet {
				defaultSuggestionsSet = true
				line.SetCompleter(defaultSuggest)
			}

			t, err := line.Prompt("> ")
			if err == liner.ErrPromptAborted {
				fmt.Println("Aborted")
				return
			} else if err != nil {
				log.Panic(err)
			}

			if t == "end" {
				fmt.Println("Aborted")
				break
			}

			switch t {
			case "cmd":
				// Command Mode
				defaultSuggestionsSet = false
				line.SetCompleter(commandSuggest)

				cli.CommandMode(func(prompt string) string {
					t, err := line.Prompt(prompt)
					if err != nil {
						log.Panic(err)
					}
					return t
				}, func(cmd string) {
					line.AppendHistory(strings.TrimSpace(cmd))
				}, prsr)

			case "faker":
				// Fake Data Generation Mode
				defaultSuggestionsSet = false
				line.SetCompleter(noneSuggest)
				cli.FakerMode(func(prompt string) string {
					t, err := line.Prompt(prompt)
					if err != nil {
						log.Panic(err)
					}
					return t
				})

			case "log":
				// Log File Config Mode
				defaultSuggestionsSet = false
				line.SetCompleter(noneSuggest)
				closer := cli.LogMode(func(prompt string) string {
					t, err := line.Prompt(prompt)
					if err != nil {
						log.Panic(err)
					}
					return t
				})
				closeFuncs = append(closeFuncs, closer)

			default:
				log.Panic("Wrong Command!!")
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
