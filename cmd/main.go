package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"natsp/pkg/connect"
	"natsp/pkg/figlet"
	"natsp/pkg/file"

	"github.com/c-bata/go-prompt"
)

var nc *nats.EncodedConn

func executor(in string) {
	fmt.Println("Your input: " + in)
		switch in {
		case "Struct":
			file.GenerateStruct()
		case "Connect":
			nc = connect.NewConnection()
		case "Publish":
			//connect.Publish()
		case "Subscribe":
		case "Help":
		case "About":
		default:
			log.Fatal("method not found !")
		}
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Struct", Description: "Add your struct file"},
		{Text: "Connect", Description: "Connect to NATS"},
		{Text: "Publish", Description: "Publish with your struct"},
		{Text: "Subscribe", Description: "Subscribe with your struct"},
		{Text: "Help", Description: "Help for work with NATS.P"},
		{Text: "About", Description: "About NATS.P"},
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}


func main() {
	print("\033[H\033[2J")
	figlet.ShowFiglet("NATS.P")
	fmt.Println("Select Command:")

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("NATS.P: A Simple NATS Prompt"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.Black))
	p.Run()
}
