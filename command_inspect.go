package main

import "fmt"

func commandInspect(cfg *config, arguments ...string) error {
	name := arguments[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")

	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
