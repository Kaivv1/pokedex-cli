package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}
