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


func GetLocationBodyFromUrl(apiURL string) ([]byte, error) {
	url := BaseURL + "location-area/"

	if apiURL != "" {
		url = apiURL
	}

	body, err := GetBodyFromUrl(url)

    if err != nil {
        return nil, err
    }

	return body, nil
}

func GetLocationsFromUrl(apiURL *string) (RespLocationsList, error) {
    body, err := GetLocationBodyFromUrl(*apiURL)

    if err != nil {
        return RespLocationsList{}, err
    }

	return GetLocationsFromBody(body), nil
}

func GetLocationsFromBody(body []byte) RespLocationsList {
	var locations RespLocationsList

	if err := json.Unmarshal(body, &locations); err != nil {
		log.Println("Couldn't unmarshall json of locations list.")
		log.Fatal(err)
	}

	return locations
}
