package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func commandMap(cfg *Config, args ...string) error {
	locationAreas, err := cfg.Pokeapi.GetLocationAreas(cfg.NextLocationUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		str := fmt.Sprintf(" - %s", area.Name)
		fmt.Println(str)
	}
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
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		str := fmt.Sprintf(" - %s", area.Name)
		fmt.Println(str)
	}
	cfg.NextLocationUrl = locationAreas.Next
	cfg.PreviousLocationUrl = locationAreas.Previous

	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	areaPokemons, err := cfg.Pokeapi.GetAreaPokemons(args[1])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[1])
	time.Sleep(time.Second)
	fmt.Println("Found pokemon: ")
	for _, pokemon := range areaPokemons.PokemonEncounters {
		str := fmt.Sprintf(" - %s", pokemon.Pokemon.Name)
		fmt.Println(str)

	}
	return nil
}

var failed bool

func commandCatch(cfg *Config, args ...string) error {
	pokemon, err := cfg.Pokeapi.GetPokemonInformation(args[1])
	if err != nil {
		return err
	}
	if _, exists := cfg.Pokedex.Get(pokemon.Name); exists {
		return fmt.Errorf("you already have a %s and it's angry that you want to replace him >_<", pokemon.Name)
	}

	if failed {
		fmt.Println("You stand there, desperate 0_0")
	} else {
		fmt.Println("You are sneaking...")
	}
	time.Sleep(time.Second)
	if failed {
		fmt.Printf("Throwing a Pokeball at %s again...\n", args[1])
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", args[1])
	}
	time.Sleep(time.Second)
	fmt.Printf("You hit %s's head...\n", pokemon.Name)
	time.Sleep(time.Second)
	threshold := 50
	randomRoll := rand.Intn(pokemon.BaseExperience)
	if randomRoll < threshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		failed = true
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Pokedex.Add(pokemon)
		failed = false
	}

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
