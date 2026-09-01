package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		inputSlice := cleanInput(input)

		command, ok := CommandRegistry[inputSlice[0]]

		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(config)

			if err != nil {
				fmt.Println(fmt.Errorf("Error running command '%s': %v", command.name, err))
			}
		}
	}
}

func cleanInput(input string) []string {
	splitString := []string{}
	input = strings.ToLower(input)
	for _, word := range strings.Split(input, " ") {
		if word != "" {
			splitString = append(splitString, word)
		}
	}

	return splitString
}
