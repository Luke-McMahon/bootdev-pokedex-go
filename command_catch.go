package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *commandConfig, args ...string) error {
	if len(args) < 1 {
		return errors.New("must provide pokemon to catch")
	}

	name := args[0]
	pokemon, err := config.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokemon at %s...\n", pokemon.Name)
	if chance > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	config.Pokedex[pokemon.Name] = pokemon
	return nil
}
