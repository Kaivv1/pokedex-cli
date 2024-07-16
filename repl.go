package main

import (
	"fmt"
	"log"

	"github.com/chzyer/readline"
)

func getCommands() map[string]ClientCommand {
	return map[string]ClientCommand{
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
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
			command.callback()
		} else {
			fmt.Println("There is no such command: " + input)
		}
	}
}
