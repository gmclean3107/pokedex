package main

import "fmt"

func commandHelp(config *config) error {

	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range config.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
