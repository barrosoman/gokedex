package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Location []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getBodyFromUrl(url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func returnLocations(body []byte) Locations {
	var locations Locations

	err := json.Unmarshal(body, &locations)

	if err != nil {
		log.Fatal(err)
	}

	return locations
}
