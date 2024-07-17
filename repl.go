package main

import (
	"fmt"
	"log"

	types "github.com/Kaivv1/pokedex-cli/types"
	"github.com/chzyer/readline"
)

func getCommands() map[string]types.ClientCommand {
	return map[string]types.ClientCommand{
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

func startRepl() {
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
			command.Callback()
		} else {
			fmt.Println("There is no such command: " + input)
		}
	}
}
