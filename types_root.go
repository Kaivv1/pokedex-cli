package main

import (
	"github.com/Kaivv1/pokedex-cli/internal/api"
	"github.com/Kaivv1/pokedex-cli/internal/pokedex"
)

type ClientCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	NextLocationUrl     *string
	PreviousLocationUrl *string
	Pokeapi             *api.Client
	Pokedex             *pokedex.Pokedex
}
