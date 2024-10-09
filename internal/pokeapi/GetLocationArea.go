package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetLocationArea(locationArea string) (LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationArea)

	// Check cache first
	if cachedData, ok := c.cache.Get(url); ok {
		var locationDetails LocationArea
		err := json.Unmarshal(cachedData, &locationDetails)
		if err != nil {
			return LocationArea{}, err
		}
		return locationDetails, nil
	}

	// If not in cache, make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Parse the JSON into a LocationArea struct
	var locationDetails LocationArea
	err = json.Unmarshal(body, &locationDetails)
	if err != nil {
		return LocationArea{}, err
	}

	// Store in cache
	c.cache.Add(url, body)

	return locationDetails, nil
}

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
