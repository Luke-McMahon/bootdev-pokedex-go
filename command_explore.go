package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *commandConfig, args ...string) error {
	if len(args) < 1 {
		return errors.New("must provide location to explore")
	}
	locationName := args[0]
	location, err := config.Client.ListLocationDetails(locationName)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", location.Name)
	fmt.Printf("Found Pokemon:\n")
	for _, detail := range location.PokemonEncounters {
		fmt.Printf("\t- %v\n", detail.Pokemon.Name)
	}
	return nil
}
