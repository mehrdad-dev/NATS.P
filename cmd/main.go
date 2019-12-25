package main

import (
	"fmt"
	"nats/completer"
	"github.com/c-bata/go-prompt"
)

func select_function(function string) {
	switch function {
	case "Models":
	case "Publish":
	case "Subscribe":
	case "Help":
	case "About":
	default:
		log.Fatal("method not found !")
		completer.terminal()
	}
}

func main() {
	fmt.Println("Please select table.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}
