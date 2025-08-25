# CO₂ Backend Code Walkthrough

This document provides an overview of the CO₂ backend implementation focusing on the in-memory caching layer.

## Connector: `internal/connectors/co2_noaa.go`

- Fetches NOAA daily CO₂ CSV data using HTTP with retry and exponential backoff.
- Parses CSV into a slice of `Point` structs containing time and value.
- Skips missing or flagged rows and drops NaN values.
- Returns the latest data point and the full slice.

## Cache Layer: `internal/cache/memory.go` and `internal/connectors/co2_cache.go`

- `memory.go` implements a thread-safe in-memory cache with TTL support.
- Cache entries expire after a configurable TTL (default 30 minutes).
- `co2_cache.go` provides `FetchCO2DataCached` which wraps the connector fetch function.
- Uses cache keys `co2:latest` and `co2:daily:30d` to store latest and 30-day data.
- On cache miss or expiration, fetches fresh data and updates cache.
- On fetch failure, returns cached data if available as fallback.

## Testing

- Unit tests for cache set/get and TTL behavior.
- Integration test for cache wrapper using a local sample CSV file.
- Tests ensure cache correctness and fallback behavior.

This caching strategy improves performance by reducing repeated upstream calls and enhances resilience by providing cached data on failures.
