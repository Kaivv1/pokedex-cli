package main

import "github.com/Kaivv1/pokedex-cli/internal/api"

type ClientCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	NextLocationUrl     *string
	PreviousLocationUrl *string
	Pokeapi             api.Client
}
