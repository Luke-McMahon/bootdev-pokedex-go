package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PokeAPI() {
	fmt.Println("Hello from PokeAPI")
}

type locationAPIResults struct {
	Count     int        `json:"count"`
	Next      string     `json:"next"`
	Previous  string     `json:"previous"`
	Locations []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations(url string) (locationAPIResults, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	resp, err := http.Get(url)
	if err != nil {
		return locationAPIResults{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAPIResults{}, err
	}

	locations := make([]location, 20)

	var response locationAPIResults
	jsonErr := json.Unmarshal(body, &response)

	if jsonErr != nil {
		return response, jsonErr
	}
	for i, location := range response.Locations {
		locations[i] = location
	}

	return response, err
}
