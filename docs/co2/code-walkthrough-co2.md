# CO₂ Tile Code Walkthrough

This document provides an overview of the CO₂ tile implementation in the One Earth project, covering both backend and frontend components.

---

## Backend

- **Connector:** `internal/connectors/co2_noaa.go`  
  Fetches NOAA daily CO₂ CSV data with retry/backoff, parses into time-value points, filters invalid data, and caches results in-memory with TTL.

- **Cache Wrapper:** `internal/connectors/co2_cache.go`  
  Provides cached access to latest and 30-day CO₂ data slices, falling back to cache on upstream failure.

- **Handlers:** `internal/http/handlers/co2.go`  
  Exposes REST API endpoints:

  - `GET /api/metrics/co2` returns latest CO₂ metric.
  - `GET /api/series/co2?days=30` returns last 30 days CO₂ series.

- **Routing:** Wired in `internal/http/router.go`.

- **Error Handling:** Returns cached data on upstream failure; returns 503 with error JSON if no cache.

- **Tests:** Cover CSV parsing, cache behavior, and API responses.

---

## Frontend

- **Schemas:** `src/schemas/co2.ts`  
  Zod schemas for CO₂ metric and series data, enabling type-safe validation.

- **Hooks:** `src/hooks/useCO2.ts`  
  React Query hooks:

  - `useCO2Latest()` fetches latest CO₂ metric.
  - `useCO2Series(days)` fetches CO₂ time series.

- **Tile Component:** `src/tiles/CO2Tile.tsx`

  - Displays last updated timestamp (UTC).
  - Shows latest CO₂ ppm value (rounded to 0.1).
  - Renders sparkline chart for last 30 days using Vega-Lite.
  - Shows muted source/license text.
  - Handles loading and error states gracefully.

- **Integration:**  
  CO2Tile is integrated into the main app grid in `src/App.tsx`.

- **React Query Setup:**  
  `QueryClientProvider` wraps the app in `src/main.tsx` to provide React Query context.

---

## Notes

- The CO2 tile fetches data from the backend API endpoints and updates reactively.
- Vega-Lite spec is defined in `src/charts/co2.spec.ts` for consistent chart styling.
- Environment variables are accessed safely with optional chaining in config files.
- Proxy configuration in `vite.config.ts` forwards API requests to backend port 8080.

---

## Next Steps

- Confirm backend Dockerfile builds with new code.
- Update frontend `.env` or API base util if needed.
- Validate CORS settings for dev and production.
- Ensure API responses pass Zod validation.
- Perform visual QA for sparkline rendering and theming.
