package cache

import (
	"sync"
	"time"
)

// CacheItem represents a single cache entry
type CacheItem struct {
	value      interface{}
	expiration int64
}

// MemoryCache is a simple in-memory cache with TTL support
type MemoryCache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
	ttl   time.Duration
}

// NewMemoryCache creates a new MemoryCache with the given TTL
func NewMemoryCache(ttl time.Duration) *MemoryCache {
	return &MemoryCache{
		items: make(map[string]CacheItem),
		ttl:   ttl,
	}
}

// SetTTL updates the TTL duration for the cache
func (c *MemoryCache) SetTTL(ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ttl = ttl
}

// Set adds a value to the cache with the default TTL
func (c *MemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(c.ttl).UnixNano(),
	}
}

// Get retrieves a value from the cache if it exists and is not expired
func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if time.Now().UnixNano() > item.expiration {
		// Item expired, delete it
		c.mu.RUnlock()
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		c.mu.RLock()
		return nil, false
	}
	return item.value, true
}

// Delete removes a key from the cache
func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
