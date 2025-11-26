package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range getCommands() {
		msg := fmt.Sprintf("%s: %s", cmd.name, cmd.description)
		fmt.Println(msg)
	}
	fmt.Println()

	return nil
}
