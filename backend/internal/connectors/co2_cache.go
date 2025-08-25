package connectors

import (
	"context"
	"time"

	"one-earth-api/internal/cache"
)

const (
	cacheKeyLatest  = "co2:latest"
	cacheKeyDaily30 = "co2:daily:30d"
)

var memCache = cache.NewMemoryCache(30 * time.Minute)

// FetchCO2DataCached fetches CO2 data with in-memory caching.
// It returns cached data if available and valid, otherwise fetches fresh data and updates cache.
func FetchCO2DataCached(ctx context.Context, url string) (latest Point, all []Point, err error) {
	// Try to get latest from cache
	cachedLatest, foundLatest := memCache.Get(cacheKeyLatest)
	cachedDaily, foundDaily := memCache.Get(cacheKeyDaily30)

	if foundLatest && foundDaily {
		latest, _ = cachedLatest.(Point)
		all, _ = cachedDaily.([]Point)
		return latest, all, nil
	}

	// Cache miss or expired, fetch fresh data
	latest, all, err = FetchCO2Data(ctx, url)
	if err != nil {
		// If fetch fails but cache exists, return cached data as fallback
		if foundLatest && foundDaily {
			latest, _ = cachedLatest.(Point)
			all, _ = cachedDaily.([]Point)
			err = nil
			return latest, all, nil
		}
		return
	}

	// Update cache
	memCache.Set(cacheKeyLatest, latest)
	memCache.Set(cacheKeyDaily30, all)

	return latest, all, nil
}
