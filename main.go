package main

import (
	"time"

	"github.com/Kaivv1/pokedex-cli/internal/api"
	"github.com/Kaivv1/pokedex-cli/internal/pokedex"
)

func main() {
	config := Config{
		Pokeapi: api.NewClient(time.Minute * 20),
		Pokedex: pokedex.NewPokedex(),
	}

	startRepl(&config)
}
