package main

import (
	//"encoding/json"
	//"fmt"
	//"io"
	//"net/http"
	"errors"
	"fmt"
)

func commandMapf(cfg *config, _ []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, _ []string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

/*

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}
	return getAndPrintLocations(cfg, url)
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	return getAndPrintLocations(cfg, *cfg.prevLocationsURL)
}

func getAndPrintLocations(cfg *config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	locationsResp := ShallowLocationsResp{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

type ShallowLocationsResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
*/
