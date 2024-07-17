package cache

import "time"

type Cache struct {
	Cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}
