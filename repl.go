package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

func getArgs(input string) ([]string, error) {
	lowerCaseInput := strings.ToLower(input)
	args := strings.Fields(lowerCaseInput)
}

func getCommands() map[string]ClientCommand {
	return map[string]ClientCommand{
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "See a list of all the Pokemon in a given area",
			Callback:    commandExplore,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}
}

func startRepl(cfg *Config) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "pokedex > ",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer rl.Close()

	for {
		input, err := rl.Readline()
		if err != nil {
			log.Fatal(err)
			break
		}

		commands := getCommands()
		if command, exists := commands[input]; exists {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("There is no such command: " + input)
		}
	}
}
