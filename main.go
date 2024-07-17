package main

import "github.com/Kaivv1/pokedex-cli/internal/api"

func main() {
	config := Config{
		Pokeapi: api.NewClient(),
	}

	startRepl(&config)
}
