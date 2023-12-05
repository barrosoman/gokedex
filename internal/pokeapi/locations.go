package pokeapi

import (
	"encoding/json"
	"log"
)

type RespLocationsList struct {
	Count     int `json:"count"`
	Next      any `json:"next"`
	Previous  any `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func GetLocationBodyFromUrl(apiURL string) []byte {
	url := baseURL + "location-area/"

	if apiURL != "" {
		url = apiURL
	}

	body := getBodyFromUrl(url)

	return body
}

func GetLocationsFromUrl(apiURL *string) RespLocationsList {
    body := GetLocationBodyFromUrl(*apiURL)

	return GetLocationsFromBody(body)
}

func GetLocationsFromBody(body []byte) RespLocationsList {
	var locations RespLocationsList

	if err := json.Unmarshal(body, &locations); err != nil {
		log.Println("Couldn't unmarshall json of locations list.")
		log.Fatal(err)
	}

	return locations
}
