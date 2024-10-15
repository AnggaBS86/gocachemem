package main

import (
	"sync"
	"testing"
	"time"
)

func TestCache_Set(t *testing.T) {
	type fields struct {
		items map[string]CacheData[any]
		mu    sync.RWMutex
	}

	mut := sync.RWMutex{}
	items := map[string]CacheData[any]{
		"key1": CacheData[any]{Value: "test"},
		"key2": CacheData[any]{Value: "test2"},
	}

	type args struct {
		key      string
		value    interface{}
		duration time.Duration
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Test case #1", fields{items, mut}, args{"foo", "bar", 5 * time.Second}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := &Cache{
				items: tt.fields.items,
				mu:    tt.fields.mu,
			}
			cache.Set(tt.args.key, tt.args.value, tt.args.duration)
		})
	}
}

func TestMain(t *testing.T) {
	cache := NewMemCache()

	// Set cache values
	cache.Set("foo", "bar", 3*time.Second)
	cache.Set("baz", 123, 5*time.Second)

	// Get values from the cache
	value, found := cache.Get("foo")

	//case not found
	cache.Get("foo1")

	time.Sleep(2 * time.Second) // Wait for "foo" to expire

	value, found = cache.Get("foo")
	if value == "bar" && found == true {
		t.Log("test successfull")
	}

	cache.Delete("foo")

	cache.DeleteExpiredCache()
}

func TestMainExpired(t *testing.T) {
	cache := NewMemCache()

	// Set cache values
	cache.Set("foo", "bar", 3*time.Second)
	cache.Set("baz", 123, 5*time.Second)

	// Get values from the cache
	value, found := cache.Get("foo")

	time.Sleep(1 * time.Second) // Wait for "foo" to expire

	value, found = cache.Get("foo")
	if value == "bar" && found == true {
		t.Log("test successfull")
	}

	cache.Set("baz1", 123, 1*time.Second)

	cache.Delete("foo")

	cache.DeleteExpiredCache()
}
