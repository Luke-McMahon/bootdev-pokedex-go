package pokeapi

import (
	"encoding/json"
	"io"
	"errors"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResponse := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResponse := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResponse, nil
}

func (c *Client) ListLocationDetails(locationName *string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/"
	if locationName == nil {
		return RespDeepLocations{}, errors.New("must provide a location to explore")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}

	locationDetails := RespDeepLocations{}
	err = json.Unmarshal(dat, &locationDetails)
	if err != nil {
		return RespDeepLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationDetails, nil
}







