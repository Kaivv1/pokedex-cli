package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetLocationAreas(url *string) (LocationArea, error) {
	fullUrl := baseUrl + "/location-area"
	if url != nil {
		fullUrl = *url
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
	return locationArea, nil
}
