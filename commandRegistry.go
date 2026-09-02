package main

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
	"map": {
		name:        "map",
		description: "Displays next 20 location areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous 20 location areas",
		callback:    commandMapB,
	},
}
