package main

import (
	"fmt"
	"os"
)

func commandMap() error {
	return nil
}

func commandMapb() error {
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for key := range commands {
		command := commands[key]
		str := fmt.Sprintf("%s: %s", command.name, command.description)
		fmt.Println(str)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
