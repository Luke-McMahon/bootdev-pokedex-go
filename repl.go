package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {
	input = strings.Trim(input, " ")
	if len(input) == 0 {
		return []string{}
	}
	input = strings.ToLower(input)
	cleaned := strings.Split(input, " ")
	return cleaned
}

type commandConfig struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig) error
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
			description: "Displays a list of 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays a list of 20 locations",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	config := commandConfig{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		command, ok := commands[text]
		if !ok {
			fmt.Printf("%v is not supported, see help for usage.\n", text)
			continue
		}

		err := command.callback(&config)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
	}
}
