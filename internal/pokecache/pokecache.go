package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
	done     chan struct{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false

}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		mutex:    sync.Mutex{},
		interval: interval,
		done:     make(chan struct{}),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			now := time.Now()
			for key, entry := range c.entries {
				if now.Sub(entry.createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		case <-c.done:
			return
		}
	}
}

func (c *Cache) Close() {
	close(c.done)
}
