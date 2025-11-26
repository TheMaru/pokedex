package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arguments ...string) error {
	if len(arguments) < 1 {
		return errors.New("no location given")
	}
	locationRes, err := cfg.pokeapiClient.GetLocationAreaData(arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arguments[0])
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range locationRes.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
