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

- The correct import path format for this project is to use the project module name as the prefix.
- For example, in the backend Go code, import internal packages like this:
  `"one-earth-api/internal/connectors"`
- Use the package name declared inside the Go source files for referencing (e.g., `connectors`).
- Avoid using relative or ambiguous import paths like `"internal/connectors"` or `"backend/internal/connectors"`.

---

## Directive to Utilize Context7 MCP Server

The Context7 MCP server can be leveraged to enhance development by providing better code, library information, and tooling support for the CO₂ data backend.

- Use the Context7 MCP server to assist in managing cache retrieval and refresh operations by integrating its tools that provide authoritative, up-to-date library and code insights.
- Replace or augment the existing in-memory cache wrapper function `FetchCO2DataCached` in `internal/connectors/co2_cache.go` to utilize MCP server tools for fetching, caching, and validating CO₂ data.
- Utilize MCP server capabilities to resolve library IDs, fetch documentation, and ensure best practices in code usage and dependencies.
- Ensure fallback to cached data on upstream fetch failure remains intact.
- Configure TTL and cache keys through MCP server settings if supported.
- Update backend handlers to interact with MCP server endpoints or tools as needed.
- Maintain existing import path conventions and package references.

---

## Frontend Schema Definitions Added

- Added Zod schema definitions for CO₂ data in `frontend/src/schemas/co2.ts`.
- Schemas include:
  - `PointSchema` for individual CO₂ data points with timestamp `T` and value `V`.
  - `MetricLatestSchema` for latest CO₂ metric with `timestamp` and `value`.
  - `SeriesSchema` as an array of `PointSchema`.
- Installed `zod` package in frontend dependencies.
- Fixed TypeScript typing issues in schema file.

This update enables type-safe validation and parsing of CO₂ data fetched from backend API endpoints in the frontend React application.
