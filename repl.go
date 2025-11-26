package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TheMaru/pokedexcli/internal/pokeapi"
	"github.com/TheMaru/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeapiCache     *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		arguments := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			if err := command.callback(cfg, arguments...); err != nil {
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
		"explore": {
			name:        "explore",
			description: "Display pokemon of a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
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
