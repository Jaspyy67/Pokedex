package main

import (
	"fmt"
	"time"

	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a pokeball at %s", pokemonName)
	fmt.Println()

	rand.Seed(time.Now().UnixNano())

	CaptureChance := 100 - pokemon.BaseExperience/4
	if CaptureChance < 10 {
		CaptureChance = 10
	}

	roll := rand.Intn(100)

	if roll <= CaptureChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Printf("You may now inspect it with the inspect command.\n\n")
		cfg.caughtPokemon[pokemonName] = *pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		fmt.Println()
	}

	return nil
}
