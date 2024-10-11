package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetPokemon(name string) (*Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}
