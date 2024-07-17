package main

import (
	"time"

	"github.com/Kaivv1/pokedex-cli/internal/api"
)

func main() {
	config := Config{
		Pokeapi: api.NewClient(time.Minute * 20),
	}

	startRepl(&config)
}
