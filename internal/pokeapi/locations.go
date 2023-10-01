package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func getBodyFromUrl(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Couldn't get response from URL \"%s\".\n", url)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Couldn't read body of http response.")
		log.Fatal(err)
	}

	return body
}

func GetLocationBodyFromUrl(apiURL string) []byte {
	url := baseURL + "location/"

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
