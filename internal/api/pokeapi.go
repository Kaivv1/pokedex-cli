package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetLocationAreas() {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatalf("Reading body gave error: %v", err)
	}
	fmt.Println(string(body))
}
