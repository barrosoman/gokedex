package pokeapi

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func GetBodyFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Println("Couldn't get response from URL '" + url + "'.")
		return nil, errors.New("Couldn't get response from URL '" + url + "'.")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Couldn't read body of http response.")
		return nil, errors.New("Couldn't read body of http response.")
	}

	return body, nil
}
