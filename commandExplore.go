package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *commandConfig, locationName *string) error {
	if locationName == nil {
		return errors.New("must provide location to explore")
	}

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
