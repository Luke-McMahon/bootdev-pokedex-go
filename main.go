package main

import (
	"time"

	"github.com/luke-mcmahon/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, time.Minute * 5)
	config := &commandConfig{
		Client: client,
	}
	startRepl(config)
}
