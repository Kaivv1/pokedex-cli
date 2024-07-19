package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Kaivv1/pokedex-cli/internal/cache"
	"github.com/Kaivv1/pokedex-cli/internal/pokedex"
)

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: cache.NewCache(cacheInterval),
	}
}

func (c *Client) GetLocationAreas(url *string) (LocationArea, error) {
	endPoint := "/location-area?offset=0&limit=20"
	fullUrl := baseUrl + endPoint
	if url != nil {
		fullUrl = *url
	}

	if data, ok := c.cache.Get(fullUrl); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationArea, nil
}

func (c *Client) GetAreaPokemons(areaName string) (AreaPokemons, error) {
	endPoint := "/location-area/" + areaName
	fullUrl := baseUrl + endPoint

	if data, ok := c.cache.Get(fullUrl); ok {
		areaPokemons := AreaPokemons{}
		err := json.Unmarshal(data, &areaPokemons)
		if err != nil {
			return AreaPokemons{}, err
		}
		return areaPokemons, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return AreaPokemons{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return AreaPokemons{}, err
	}

	defer res.Body.Close()
	if res.StatusCode > 399 {
		return AreaPokemons{}, fmt.Errorf("bad status code %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaPokemons{}, err
	}
	areaPokemons := AreaPokemons{}
	err = json.Unmarshal(data, &areaPokemons)
	if err != nil {
		return AreaPokemons{}, err
	}
	c.cache.Add(fullUrl, data)

	return areaPokemons, nil
}

func (c *Client) GetPokemonInformation(name string) (pokedex.Pokemon, error) {
	endpoint := "/pokemon/" + name
	fullUrl := baseUrl + endpoint

	if data, ok := c.cache.Get(fullUrl); ok {
		pokemon := pokedex.Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return pokedex.Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return pokedex.Pokemon{}, fmt.Errorf("bad status code %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	pokemon := pokedex.Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	c.cache.Add(fullUrl, data)

	return pokemon, nil
}
