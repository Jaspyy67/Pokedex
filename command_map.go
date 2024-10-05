package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokedex struct {
	Next     string
	Previous string
}

var pokedexData pokedex

func commandMap(url string) (pokedex, error) {
	resp, err := http.Get(pokeUrl)
	if err != nil {
		return pokedex{}, fmt.Errorf("failed to fetch location areas: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokedex{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var loc location
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return pokedex{}, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}

	return pokedex{
		Next:     loc.Next,
		Previous: loc.Previous,
	}, nil
}

func commandMapf() error {
	if pokedexData.Next == "" {
		return fmt.Errorf("No next page available")
	}
	var err error
	pokedexData, err = commandMap(pokedexData.Next)
	if err != nil {
		return fmt.Errorf("failed to execute mapb: %v", err)
	}
	return nil
}

func commandMapb() error {
	if pokedexData.Previous == "" {
		return fmt.Errorf("No previous page available")
	}
	var err error
	pokedexData, err = commandMap(pokedexData.Previous)
	if err != nil {
		return fmt.Errorf("failed to execute mapb: %v", err)
	}
	return nil
}
