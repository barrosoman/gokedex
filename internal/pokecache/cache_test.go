package pokecache_test

import (
	"fmt"
	"pokedex-go/internal/pokecache"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
    interval := 5 * time.Millisecond

    cases := []struct {
        key string
        val []byte
    }{
        {
            key: "www.google.com",
            val: []byte("google"),
        },
        {
            key: "www.amazon.com",
            val: []byte("amazon"),
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
    }
}

func TestCacheReap(t *testing.T) {
    baseDuration := 5 * time.Millisecond
    waitDuration := baseDuration * 2

    url := "www.cache.org"
    val := []byte("cache")

    cache := pokecache.NewCache(baseDuration)

    cache.Add(url, val)

    _, ok := cache.Get(url)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

    time.Sleep(waitDuration)

    _, ok = cache.Get(url)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
