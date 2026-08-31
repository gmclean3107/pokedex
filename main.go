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

		fmt.Printf("Your command was: %s\n", inputSlice[0])
	}
}
