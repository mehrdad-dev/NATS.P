package completer

import (
	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Models", Description: "Manage Models"},
		{Text: "Publish", Description: "Publish With Models"},
		{Text: "Subscribe", Description: "Subscribe With Models"},
		{Text: "Help", Description: "Help for Work With NATS.P"},
		{Text: "About", Description: "About for NATS.P"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func terminal () string {
	var d prompt.Document
	p := completer(d)
	p = p

	terminal := prompt.Input(">>> ", completer,
		prompt.OptionTitle("nats-prompt"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.Green),
		prompt.OptionSuggestionBGColor(prompt.Black))

	return terminal
}