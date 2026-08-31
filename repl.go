package main

import "strings"

func cleanInput(input string) []string {
	splitString := []string{}

	for _, word := range strings.Split(input, " ") {
		if word != "" {
			splitString = append(splitString, word)
		}
	}

	return splitString
}
