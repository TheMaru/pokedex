package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		commands := cleanInput(scanner.Text())
		returnMsg := fmt.Sprintf("Your command was: %s", commands[0])
		fmt.Println(returnMsg)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
	}
}
