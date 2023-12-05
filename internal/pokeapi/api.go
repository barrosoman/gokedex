package pokeapi

import (
	"io"
	"log"
	"net/http"
)

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
