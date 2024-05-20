package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry

	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	result := Cache{
		entries:  make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go result.reapLoop(interval)

	return result
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.removeOldEntries(time.Now().UTC(), interval)
	}
}

func (c *Cache) removeOldEntries(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
