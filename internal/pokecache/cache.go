package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type Cache struct {
    entries map[string] cacheEntry
    mu *sync.Mutex
}

func NewCache(timeout time.Duration) Cache {
    newCache := Cache{entries : make(map[string]cacheEntry), mu : &sync.Mutex{}}
    go newCache.reapLoop(timeout)
    return newCache
}

func (c *Cache) reapLoop(timeout time.Duration) {
    tick := time.NewTicker(timeout)
    for range tick.C {
        c.reap(time.Now().UTC(), timeout)
    }
}

func (c *Cache) reap(now time.Time, timeout time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    for k, v := range c.entries {
        if v.createdAt.Before(now.Add(-timeout)) {
            delete(c.entries, k)
        }
    }
}

func (c Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}

}

func (c Cache) Entries() map[string]cacheEntry {
    return c.entries
}

func (c Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    entry, ok := c.entries[key]

    if !ok {
        return nil, ok
    }

    val := entry.val

    return val, ok
}

func (entry cacheEntry) Val() []byte {
    return entry.val
}
