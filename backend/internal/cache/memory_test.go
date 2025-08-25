package cache

import (
	"testing"
	"time"
)

func TestMemoryCache_SetGet(t *testing.T) {
	cache := NewMemoryCache(50 * time.Millisecond)

	cache.Set("key1", "value1")
	val, found := cache.Get("key1")
	if !found {
		t.Fatal("Expected to find key1")
	}
	if val.(string) != "value1" {
		t.Fatalf("Expected value1, got %v", val)
	}

	// Wait for expiration
	time.Sleep(60 * time.Millisecond)
	_, found = cache.Get("key1")
	if found {
		t.Fatal("Expected key1 to be expired")
	}
}

func TestMemoryCache_SetTTL(t *testing.T) {
	cache := NewMemoryCache(100 * time.Millisecond)
	cache.Set("key2", "value2")

	// Wait less than TTL, key should be found
	time.Sleep(50 * time.Millisecond)
	_, found := cache.Get("key2")
	if !found {
		t.Fatal("Expected key2 to be found before TTL expiration")
	}

	// Change TTL to shorter duration
	cache.SetTTL(10 * time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	_, found = cache.Get("key2")
	if found {
		t.Fatal("Expected key2 to be expired after TTL change")
	}
}
