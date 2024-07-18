package main

import (
	"errors"
	"fmt"
	"os"
)

func commandMap(cfg *Config, args ...string) error {
	locationAreas, err := cfg.Pokeapi.GetLocationAreas(cfg.NextLocationUrl)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		str := fmt.Sprintf("- %s", area.Name)
		fmt.Println(str)
	}
	fmt.Println()
	cfg.NextLocationUrl = locationAreas.Next
	cfg.PreviousLocationUrl = locationAreas.Previous

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.PreviousLocationUrl == nil {
		return errors.New("you are on page 1")
	}
	locationAreas, err := cfg.Pokeapi.GetLocationAreas(cfg.PreviousLocationUrl)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		str := fmt.Sprintf("- %s", area.Name)
		fmt.Println(str)
	}
	fmt.Println()
	cfg.NextLocationUrl = locationAreas.Next
	cfg.PreviousLocationUrl = locationAreas.Previous

	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	areaPokemons, err := cfg.Pokeapi.GetAreaPokemons(args[1])
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Exploring %s...\n", args[1])
	fmt.Println("Found pokemon: ")
	for _, pokemon := range areaPokemons.PokemonEncounters {
		str := fmt.Sprintf("- %s", pokemon.Pokemon.Name)
		fmt.Println(str)

	}
	fmt.Println()
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
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

func commandExit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}
