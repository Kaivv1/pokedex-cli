package cache

import "time"

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.Cache[key] = cacheEntry{
		value:     value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.Cache[key]
	return cacheEntry.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	expiration := time.Now().Add(-interval)
	for k, v := range c.Cache {
		if v.createdAt.Before(expiration) {
			delete(c.Cache, k)
		}
	}
}
