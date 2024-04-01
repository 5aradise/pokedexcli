package pokecache

import (
	"time"
)

type Cache struct {
	cache     map[string]cacheEntry
}

type cacheEntry struct {
	value     []byte
	createdAt time.Time
}

func NewCache(ttl time.Duration, chunkSize int) Cache {
	c := Cache{
		cache:     make(map[string]cacheEntry),
	}
	go c.reapLoop(ttl)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{
		value:     value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, isExist := c.cache[key]
	return entry.value, isExist
}

func (c *Cache) reapLoop(interval time.Duration) {
	tiker := time.NewTicker(interval)
	for now := range tiker.C {
		delayToDelete := now.UTC().Add(-interval)
		for k, v := range c.cache {
			if v.createdAt.Before(delayToDelete) {
				delete(c.cache, k)
			}
		}
	}
}
