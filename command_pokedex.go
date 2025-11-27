package main

import "fmt"

func commandPokedex(cfg *config, arguments ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught a pokemon so far. You can do so with the catch command!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
