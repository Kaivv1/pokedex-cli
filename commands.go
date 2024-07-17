package main

import (
	"fmt"
	"os"

	"github.com/Kaivv1/pokedex-cli/internal/api"
)

func commandMap() error {
	areas, _ := api.GetLocationAreas("next")
	for _, area := range areas {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb() error {
	areas, err := api.GetLocationAreas("previous")
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, area := range areas {
		fmt.Println(area.Name)
	}
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
		str := fmt.Sprintf("%s: %s", command.Name, command.Description)
		fmt.Println(str)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
