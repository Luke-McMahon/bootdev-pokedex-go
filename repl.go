package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/luke-mcmahon/pokedexcli/internal/pokeapi"
)

type commandConfig struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig, *string) error
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
			name: "explore",
			description: "Explore the passed in location",
			callback: commandExplore,
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
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		command, ok := commands[cleaned[0]]
		if !ok {
			fmt.Printf("%v is not supported, see help for usage.\n", text)
			continue
		}

		var loc string
		if len(cleaned) > 1 {
			loc = cleaned[1]
		}
		err := command.callback(config, &loc)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}
