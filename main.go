package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			err := command.callback(CommandRegistry)

			if err != nil {
				fmt.Println(fmt.Errorf("Error running command '%s': %v", command.name, err))
			}
		}
	}
}
