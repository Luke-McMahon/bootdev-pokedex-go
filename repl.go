package main

import (
	"bufio"
	"fmt"
	"github.com/luke-mcmahon/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type commandConfig struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Pokedex  map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the passed in location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the passed in Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(input string) []string {
	input = strings.Trim(input, " ")
	if len(input) == 0 {
		return []string{}
	}
	input = strings.ToLower(input)
	cleaned := strings.Split(input, " ")
	return cleaned
}

func startRepl(config *commandConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Printf("%v is not supported, see help for usage.\n", commandName)
			continue
		}

		err := command.callback(config, args...)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}
