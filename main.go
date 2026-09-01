package main

func main() {
	config := &Config{commandRegistry: CommandRegistry}

	startRepl(config)
}
