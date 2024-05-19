package main

import (
	"errors"
	"fmt"

	"github.com/luke-mcmahon/pokedexcli/internal/pokeapi"
)

func commandMapB(config *commandConfig) error {
	if config.Previous == "" {
		return errors.New(fmt.Sprintf("cannot move further back\n"))
	}
	response, err := pokeapi.GetLocations(config.Previous)
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
