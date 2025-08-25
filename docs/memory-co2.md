# CO₂ In-Memory Cache

This document describes the in-memory caching strategy implemented for the CO₂ data backend.

- Cache is implemented in Go in `internal/cache/memory.go`.
- Supports TTL (time-to-live) for cache entries, defaulting to 30 minutes.
- Cache keys used:
  - `co2:latest` for the latest CO₂ data point.
  - `co2:daily:30d` for the last 30 days of daily CO₂ data.
- Cache wrapper function `FetchCO2DataCached` in `internal/connectors/co2_cache.go` handles cache retrieval and refresh.
- On cache miss or expiration, fresh data is fetched from NOAA and cache is updated.
- If upstream fetch fails but cache exists, cached data is returned as fallback.
- TTL is configurable via a setter method on the cache instance.
- Tests cover cache set/get, TTL expiration, and cache wrapper functionality.

This caching layer improves performance and resilience by reducing repeated upstream data fetches and providing fallback data on failures.
