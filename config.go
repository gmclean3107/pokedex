package main

type Config struct {
	commandRegistry map[string]cliCommand
}

var config Config = Config{commandRegistry: CommandRegistry}
