package main

import (
	"pokedex-go/internal/pokeapi"
)

type config struct {
    client               pokeapi.Client
    pokemons             map[string]pokeapi.Pokemon
	previousLocationsURL string
	nextLocationsURL     string
}

func (cfg *config) setPreviousLocationsURL(newURL string) {
	cfg.previousLocationsURL = newURL
}

func (cfg *config) setNextLocationsURL(newURL string) {
	cfg.nextLocationsURL = newURL
}

func (cfg *config) setupLocationURLs(locationList pokeapi.RespLocationsList) {
	newPreviousURL, _ := locationList.Previous.(string)
	newNextURL, _ := locationList.Next.(string)

	cfg.setPreviousLocationsURL(newPreviousURL)
	cfg.setNextLocationsURL(newNextURL)
}
