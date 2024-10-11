package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}
	} else {
		return fmt.Errorf("pokemon has not been caught yet")
	}

	return nil

}
