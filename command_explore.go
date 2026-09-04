package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {

	if len(cfg.arguments) < 1 {
		return errors.New("you must provide a location name")
	}

	locationResp, err := cfg.pokeapiClient.ExploreLocation(cfg.arguments[0])

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", cfg.arguments[0])
	fmt.Println("Found Pokemon: ")

	for _, encounter := range locationResp.Pokemon_encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
