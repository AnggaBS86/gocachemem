package cache

import "time"

// Define CacheData for defining Value and Expiration
type CacheData[T any] struct {
	Value      T
	Expiration int64
}

// Check the CacheData is Expired
// Return bool
func (item *CacheData[T]) IsExpired() bool {
	if item.Expiration < 1 {
		return false
	}

	return time.Now().UnixNano() > item.Expiration
}
