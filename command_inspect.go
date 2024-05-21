package main

import (
	"fmt"
)

func commandInspect(config *commandConfig, args ...string) error {
	pokemon := args[0]

	caught, ok := config.Pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", caught.Name)
	fmt.Printf("Height: %v\n", caught.Height)
	fmt.Printf("Weight: %v\n", caught.Weight)

	fmt.Printf("Stats:\n")
	for i := 0; i < len(caught.Stats); i++ {
		fmt.Printf("  - %s: %v\n", caught.Stats[i].Stat.Name, caught.Stats[i].BaseStat)
	}

	fmt.Printf("Types:\n")
	for i := 0; i < len(caught.Types); i++ {
		fmt.Printf("  - %s\n", caught.Types[i].Type.Name)
	}

	return nil
}
