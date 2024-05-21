package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *commandConfig, s *string) error {
	locationsResp, err := config.Client.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationsResp.Next
	config.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(config *commandConfig, s *string) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := config.Client.ListLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationsResp.Next
	config.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
