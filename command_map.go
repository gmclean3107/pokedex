package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int            `json: "count"`
	Next     string         `json: "next"`
	Previous string         `json: "previous"`
	Results  []LocationArea `json: "results"`
}

type LocationArea struct {
	Name string `json: "name"`
	Url  string `json: "url"`
}

func commandMap(config *Config) error {
	var res *http.Response
	var err error

	if config.mapNext == "" {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
	} else {
		res, err = http.Get(config.mapNext)
	}

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(res.Body)

	var locationRes LocationAreaResponse

	err = decoder.Decode(&locationRes)

	if err != nil {
		return err
	}

	for _, result := range locationRes.Results {
		fmt.Println(result.Name)
	}

	config.mapNext = locationRes.Next
	config.mapPrev = locationRes.Previous

	return nil
}

func commandMapB(config *Config) error {
	var res *http.Response
	var err error

	if config.mapPrev == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		res, err = http.Get(config.mapPrev)
	}

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(res.Body)

	var locationRes LocationAreaResponse

	err = decoder.Decode(&locationRes)

	if err != nil {
		return err
	}

	for _, result := range locationRes.Results {
		fmt.Println(result.Name)
	}

	config.mapPrev = locationRes.Previous
	config.mapNext = locationRes.Next

	return nil
}
