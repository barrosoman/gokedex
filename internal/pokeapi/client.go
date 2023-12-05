package pokeapi

import (
	"net/http"
	"pokedex-go/internal/pokecache"
	"time"
)

type Client struct {
	Cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
    return Client {
        Cache: pokecache.NewCache(cacheInterval),
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}
