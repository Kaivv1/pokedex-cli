package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

func getArgs(input string) []string {
	lowerCaseInput := strings.ToLower(input)
	args := strings.Fields(lowerCaseInput)
	return args
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
			Name:        "explore {location-area}",
			Description: "See a list of all the Pokemon in a given area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch {pokemon}",
			Description: "Catch a Pokemon and add it to your Pokedex",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect {pokemon}",
			Description: "inspect and see details about the pokemon you caught",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "see all of the Pokemon you have caught",
			Callback:    commandPokedex,
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
		Prompt: "Pokedex > ",
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
		args := getArgs(input)

		commands := getCommands()
		if command, exists := commands[args[0]]; exists {
			err := command.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("There is no such command: " + input)
		}
	}
}
