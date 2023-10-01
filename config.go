package main

import (
	"pokedex-go/internal/pokeapi"
	"pokedex-go/internal/pokecache"
)

type config struct {
	cache                pokecache.Cache
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
