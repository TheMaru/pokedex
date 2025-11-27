package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arguments ...string) error {
	if len(arguments) < 1 {
		return errors.New("no pokemon name given")
	}
	pokemonName := arguments[0]

	pokemonData, err := cfg.pokeapiClient.GetPokemonInformation(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchChance := max(5, 100-pokemonData.BaseExperience/2)
	// fmt.Printf("Chance is %d%%\n", catchChance)
	randPercent := rand.Intn(100)
	// fmt.Printf("rand num %d\n", randPercent)

	if randPercent < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command")
		_, exists := cfg.caughtPokemon[pokemonName]
		if !exists {
			cfg.caughtPokemon[pokemonName] = pokemonData
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
