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
	cache map[string]cacheEntry
	mutex sync.Mutex
}

func NewCache(interval time.Duration) (*Cache) {
	cache := &Cache{
		cache: map[string]cacheEntry{},
	}
	
	go cache.reapLoop(interval)
	
	return cache
}

  
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mutex.Lock()
		for key, value := range c.cache {
			age := time.Since(value.createdAt)
			if age > interval {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	data, ok := c.cache[key];
	if !ok {
		return nil, false
	}
	return data.val, true
}
