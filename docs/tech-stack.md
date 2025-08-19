# One Earth – Tech Stack (no-SSR, Docker-first)

This is a lean, high-signal list of **libraries and tools** for a Vite + React frontend and a Go API backend. It favors small deps, fast DX, and “live data” tiles. Keep it minimal at first (✳️ **Essentials**) and add “Later / optional” as the product grows.

---

## Frontend (Vite + React + TypeScript)

### ✳️ Essentials
- **Vite** – bundler/dev server (already in starter)
- **React 18** – UI (already in starter)
- **TypeScript** – typing (already in starter)
- **TanStack Query** – data fetching, caching, retries, background refresh
- **Zod** – runtime validation + TS inference for API responses
- **Tailwind CSS** – utility-first styling (fast iteration)
- **shadcn/ui** (optional baseline) – accessible primitives built on Radix
- **Radix UI** – unstyled, accessible headless components (if not using shadcn)
- **react-error-boundary** – graceful error states per tile
- **clsx** / **tailwind-merge** – simple class composition

### Charts & Maps
- **Vega-Lite** + **react-vega** – declarative, consistent charts
- **MapLibre GL JS** – open-source maps (vector tiles)
- **TopoJSON** – lightweight geographic shapes (if needed)
- **date-fns** – formatting time-series axes/ticks

> Simple rule: prefer **Vega-Lite** for most charts. Use D3 only for bespoke hero visuals later.

### Networking & Utilities
- **ky** or **axios** – tiny HTTP client (or just `fetch`)
- **event-source-polyfill** – robust **SSE** support (for live tiles) if needed
- **valtio** or **zustand** – tiny global state (if you find you need it)
- **react-helmet-async** – metadata for share cards (once public)

### Testing (FE)
- **Vitest** – unit tests
- **@testing-library/react** – component tests
- **Playwright** – E2E (add later)

### Linting & DX
- **ESLint** (typescript, react, jsx-a11y) + **Prettier**
- **vite-plugin-checker** – TS + ESLint in dev overlay (optional)
- **msw** – mock API for UI dev when backend is down (optional)

### Install (FE)
```bash
# from /frontend
yarn add @tanstack/react-query zod clsx tailwind-merge
yarn add vega vega-lite react-vega
yarn add maplibre-gl topojson-client
yarn add react-error-boundary ky

# dev deps
yarn add -D eslint prettier @typescript-eslint/eslint-plugin @typescript-eslint/parser \
  eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-jsx-a11y \
  vitest @testing-library/react @testing-library/jest-dom @types/node
```

---

## Backend (Go 1.22+)

### ✳️ Essentials
- **chi** – HTTP router (already in starter)
- **cors** – CORS middleware (already in starter)
- **slog** (stdlib) – structured logging
- **pgx/v5** – Postgres driver (add when DB is added)
- **sqlc** *or* **ent** – queries/entities generation (pick one later)
- **go-playground/validator** – payload validation
- **retryablehttp** (hashicorp) – resilient HTTP pulls for upstream data
- **robfig/cron/v3** – simple scheduled fetchers (pollers for tiles)
- **SSE** – server-sent events via stdlib (no dep) for “live” push

### Storage & Caching
- **Postgres / TimescaleDB** – time-series `observations` table
- **Redis (go-redis/v9)** – ephemeral cache + job signals
- **S3-compatible** (AWS SDK v2 or **minio-go**) – raw snapshot storage

### Observability
- **OpenTelemetry** – traces/metrics/logs exporters (otlp)
- **Prometheus client_golang** – `/metrics` endpoint
- **Sentry** – error reporting (FE/BE)

### Security & Hardening
- **unrolled/secure** – security headers
- **rate limiting** – **ulule/limiter** (or implement token bucket)
- **env** – **kelseyhightower/envconfig** or **joho/godotenv** (dev only)

### Testing (BE)
- **testify** – assertions
- **gomock** – interface mocks (optional)
- **httptest** (stdlib) – handler tests

### Install (BE)
```bash
# from /backend (once you add DB/caching)
go get github.com/go-chi/chi/v5 github.com/go-chi/cors
go get github.com/hashicorp/go-retryablehttp
go get github.com/go-playground/validator/v10
go get github.com/robfig/cron/v3
go get github.com/redis/go-redis/v9
go get github.com/jackc/pgx/v5
go get github.com/prometheus/client_golang
go get go.opentelemetry.io/otel go.opentelemetry.io/otel/sdk \
       go.opentelemetry.io/otel/trace go.opentelemetry.io/otel/exporters/otlp/otlptrace
```

---

## Data Connectors (first wave)
- **NOAA GML CO₂** – daily CSV
- **NSIDC Sea Ice Index** – daily extent data
- **NASA FIRMS** – active fires (CSV/GeoJSON; NRT/URT)
- **OpenAQ** – city PM2.5/AQI
- **NASA GISTEMP v4** – monthly anomaly
- **NASA Sea Level** – monthly GMSL
- **NOAA CPC ENSO** – weekly/monthly indices

> Connector pattern: backoff + validation (Zod on FE, sanity checks in BE), append-only persistence, edge-cache 5–30 min.

---

## Minimal Feature Decisions
- **Routing**: keep it simple with file-based component routes or a tiny router (you can add TanStack Router later).
- **State**: prefer **TanStack Query**; add a tiny global state store only if you truly need app-level state.
- **Styling**: Tailwind + a small token file; only add a component lib (shadcn) if it speeds you up.
- **Charts**: standardize on Vega-Lite for consistency; isolate specs per tile in `/src/charts`.

---

## Suggested Folders
```
frontend/
  src/
    charts/           # vega-lite specs and helpers
    tiles/            # reusable tile components (AQI, CO2, etc.)
    lib/              # API client, formatters, hooks
backend/
  internal/
    http/             # handlers, SSE
    connectors/       # NOAA, NSIDC, etc.
    store/            # pgx/sqlc glue
    jobs/             # cron tasks
```

---

## Shortlist to install **right now**
- FE: `@tanstack/react-query`, `zod`, `react-error-boundary`, `vega`, `vega-lite`, `react-vega`, `maplibre-gl`
- BE: `hashicorp/go-retryablehttp`, `robfig/cron/v3`
- Next step after MVP tiles: add Postgres (`pgx`) + Redis and `sqlc` (or `ent`) when you persist observations.
