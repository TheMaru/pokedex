package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TheMaru/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			if err := command.callback(cfg); err != nil {
				fmt.Fprintln(os.Stderr, "command error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations in the (Pokemon) world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations in the (Pokemon) world",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	lowerAndTrimmed := strings.TrimSpace(lower)
	return strings.Fields(lowerAndTrimmed)
}
