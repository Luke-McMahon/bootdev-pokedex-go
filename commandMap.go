package main

import (
	"errors"
	"fmt"

	"github.com/luke-mcmahon/pokedexcli/internal/pokeapi"
)

func commandMap(config *commandConfig) error {
	response, err := pokeapi.GetLocations(config.Next)
	if err != nil {
		return errors.New(fmt.Sprintf("Error getting locations from API: %v", err))
	}

	config.Next = response.Next
	config.Previous = response.Previous

	for _, loc := range response.Locations {
		fmt.Println(loc.Name)
	}

	return nil
}
