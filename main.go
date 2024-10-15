package cache

import (
	"sync"
	"time"
)

// Define Cache struct
// for define map items and Mutex mu
type Cache struct {
	items map[string]CacheData[any]
	mu    sync.RWMutex
}

// Define new cache
func NewCacheMem() *Cache {
	return &Cache{
		items: make(map[string]CacheData[any]),
	}
}

// Set new key of Cache
// param key string
// param value interface{}
// param duration time.Duration
func (cache *Cache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	cache.mu.Lock()
	cache.items[key] = CacheData[any]{
		Value:      value,
		Expiration: expiration,
	}

	cache.mu.Unlock()
}

// Get gets a value from the cache
// By the key
// Return interface{}, bool
func (cache *Cache) Get(key string) (interface{}, bool) {
	cache.mu.RLock()
	item, found := cache.items[key]
	cache.mu.RUnlock()

	if !found || item.IsExpired() {
		return nil, false
	}

	return item.Value, true
}

// Delete deletes a value from the cache
// Param key
func (cache *Cache) Delete(key string) {
	cache.mu.Lock()
	delete(cache.items, key)
	cache.mu.Unlock()
}

// Cleanup removes expired items from the cache
// Return void
func (cache *Cache) DeleteExpiredCache() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, item := range cache.items {
		if item.IsExpired() {
			delete(cache.items, key)
		}
	}
}
