# One Earth — CO₂ Tile Implementation Plan

This TODO plan builds on the existing docs in `docs/`:

- [`tech-stack.md`](./tech-stack.md)
- [`directory-structure.md`](./directory-structure.md)
- [`requirements.md`](./requirements.md)
- [`viz-ideas.md`](./viz-ideas.md)

---

## Goal

Implement the first visualization: **Atmospheric CO₂** from NOAA GML.

- **MVP:**  
  Live tile = current ppm + 7–30 day sparkline.
- **Next:**  
  Seasonal cycle (monthly mean + band), Climate spiral, Small multiples (YoY).

---

## Backend (Go)

**Tasks**

- [ ] Create `internal/connectors/co2_noaa.go`

  - Fetch NOAA daily CSV (Mauna Loa) using `net/http` (add retry/backoff).
  - Parse into `[]Point{ T, V }`.
  - Skip missing/flagged rows, drop NaNs.
  - Return latest + full slice.

- [ ] Add in-memory cache (`internal/cache/memory.go`)

  - TTL ~30m (configurable).
  - Keys: `co2:latest`, `co2:daily:30d`.

- [ ] Create handlers `internal/http/handlers/co2.go`

  - `GET /api/metrics/co2` → `MetricLatest`.
  - `GET /api/series/co2?days=30` → `Series`.

- [ ] Wire routes in `router.go`.

- [ ] Add error handling:

  - Return cached data on upstream fail.
  - If no cache, return `{ error: … }` + 503.

- [ ] Add tests:
  - Parse sample CSV.
  - Cache returns last good value.
  - API returns sorted, valid JSON.

---

## Frontend (React + Vite)

**Tasks**

- [ ] Add schema definitions (`src/schemas/co2.ts`) with Zod.
- [ ] Add hooks (`src/hooks/useCO2.ts`) using TanStack Query.

  - `useCO2Latest()` → fetch `/api/metrics/co2`.
  - `useCO2Series(days)` → fetch `/api/series/co2`.

- [ ] Add `src/tiles/CO2Tile.tsx`

  - Show chip: last updated UTC.
  - Show stat: latest ppm (rounded 0.1).
  - Show sparkline (last 30 days) with Vega-Lite + theme.
  - Show muted source/license.

- [ ] Integrate into `App.tsx` (add tile into grid).

- [ ] Error/Loading states:
  - Loading → “Loading…”
  - Error → “Data unavailable — retrying…”

---

## Integration

**Tasks**

- [ ] Confirm backend Dockerfile builds with new code.
- [ ] Update frontend `.env` or API base util if needed (`http://localhost:8081` in dev).
- [ ] Validate CORS: dev = `*`, prod = `https://one-earth.info`.
- [ ] Ensure `/api/metrics/co2` + `/api/series/co2` pass Zod parsing.
- [ ] Visual QA: Sparkline renders, value is shown, theme applies.

---

## Next Iterations

- [ ] Add monthly mean endpoint (`/api/series/co2-monthly`) for seasonal cycle chart.
- [ ] Add anomalies endpoint for climate spiral.
- [ ] Add small multiples (per-year) view.

---

## PR Workflow

- **PR 1:** Backend connector + endpoints + tests.
- **PR 2:** Frontend schemas + hooks + tile.
- **PR 3:** Integration in App grid, QA.
- **PR 4:** Seasonal/spiral/small multiples (future).

---

## References

- [NOAA GML CO₂ data](https://gml.noaa.gov/ccgg/trends/data.html)
- [viz-ideas.md](./viz-ideas.md) → CO₂ section
- [requirements.md](./requirements.md) → API consistency rules
- [tech-stack.md](./tech-stack.md) → libraries
- [directory-structure.md](./directory-structure.md) → placement of files
