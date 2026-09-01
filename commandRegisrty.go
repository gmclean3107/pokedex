package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

var CommandRegistry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	},
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {

	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range config.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
