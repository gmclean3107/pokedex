package main

import (
	"strings"
)

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
