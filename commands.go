package main

import (
	"errors"
	"fmt"
	"os"
)

func commandMap(cfg *Config) error {
	locationAreas, err := cfg.Pokeapi.GetLocationAreas(cfg.NextLocationUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		str := fmt.Sprintf("- %s", area.Name)
		fmt.Println(str)
	}

	cfg.NextLocationUrl = locationAreas.Next
	cfg.PreviousLocationUrl = locationAreas.Previous

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.PreviousLocationUrl == nil {
		return errors.New("you are on page 1")
	}
	locationAreas, err := cfg.Pokeapi.GetLocationAreas(cfg.PreviousLocationUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		str := fmt.Sprintf("- %s", area.Name)
		fmt.Println(str)
	}

	cfg.NextLocationUrl = locationAreas.Next
	cfg.PreviousLocationUrl = locationAreas.Previous

	return nil
}

func commandExplore(cfg *Config) error {

	return nil
}

func commandHelp(cfg *Config) error {
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

func commandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}
