package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	types "github.com/Kaivv1/pokedex-cli/types"
)

var Config types.Config

func GetLocationAreas(page string) ([]types.Area, error) {
	var url string
	firstPage := false

	switch page {
	case "next":
		if Config.Next == nil {
			url = "https://pokeapi.co/api/v2/location-area"
		} else {
			url = *Config.Next
		}
	case "previous":
		if Config.Previous == nil {
			firstPage = true
		} else {
			url = *Config.Previous
		}
	}
	if firstPage {
		return nil, errors.New("you are on page 1")
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatalf("Reading body gave error: %v", err)
	}
	var locationArea types.LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		log.Fatalf("Failed to unmarshal the body: %v", err)
	}

	Config.Next = &locationArea.Next
	Config.Previous = locationArea.Previous

	return locationArea.Results, nil
}
