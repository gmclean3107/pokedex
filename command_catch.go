package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config) error {

	if len(cfg.arguments) < 1 {
		return errors.New("Must name a pokemon to catch")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.arguments[0])

	caught, err := cfg.pokeapiClient.CatchPokemon(cfg.arguments[0])

	if err != nil {
		return err
	}

	if caught {
		fmt.Printf("%s was caught!\n", cfg.arguments[0])
	} else {
		fmt.Printf("%s escaped!\n", cfg.arguments[0])
	}

	return nil
}
