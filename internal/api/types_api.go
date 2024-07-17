package api

import (
	"github.com/Kaivv1/pokedex-cli/internal/cache"
	"net/http"
)

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
