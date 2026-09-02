package main

type Config struct {
	commandRegistry map[string]cliCommand
	mapNext         string
	mapPrev         string
}

var config Config = Config{
	commandRegistry: CommandRegistry,
	mapNext:         "",
	mapPrev:         "",
}
