package main

import "github.com/c-bata/go-prompt"

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

		{Text: "distinct_count(", Description: "Add distinct count aggregator"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func noneCompleter(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}
