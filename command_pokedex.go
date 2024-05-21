package main

import "fmt"

func commandPokedex(config *commandConfig, args ...string) error {

	fmt.Println("Your Pokedex:")

	for _, v := range config.Pokedex {
		fmt.Printf("  - %s\n", v.Name)
	}

	return nil
}
