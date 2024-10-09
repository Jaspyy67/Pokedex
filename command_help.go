package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("\nWelcome to Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
