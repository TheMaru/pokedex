package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands map[string]cliCommand

func main() {
	supportedCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    printHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		commands := cleanInput(scanner.Text())
		if cmd, ok := supportedCommands[commands[0]]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Fprintln(os.Stderr, "command error:", err)
			}
		} else {
			fmt.Println("unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func printHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range supportedCommands {
		msg := fmt.Sprintf("%s: %s", cmd.name, cmd.description)
		fmt.Println(msg)
	}

	return nil
}
