package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Kaivv1/pokedex-cli/internal/cache"
)

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: cache.NewCache(cacheInterval),
	}
}

func (c *Client) GetLocationAreas(url *string) (LocationArea, error) {
	endPoint := "/location-area"
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
		fmt.Println("this is cached")
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
	fmt.Println("this is not cached")
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

	areaPokemons := AreaPokemons{}
	err = json.Unmarshal(data, &areaPokemons)
	if err != nil {
		return AreaPokemons{}, err
	}
	c.cache.Add(fullUrl, data)

	return areaPokemons, nil
}
