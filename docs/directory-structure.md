# One Earth — Directory Structure & Architecture

A lightweight, extensible layout for a **Vite + React (Yarn)** frontend and a **Go** backend. This doc explains **where things live**, **what they do**, and **guardrails** to keep the codebase tidy as it grows.

> Monorepo root is `one-earth/` (Docker-first; no SSR).

---

## 0) Repository Layout (top-level)

```
one-earth/
├─ docker-compose.yml         # Orchestrates web (nginx+Vite build) + API (Go)
├─ README.md
├─ docs/
│  ├─ sources.md              # Data source registry & licensing
│  └─ tech-stack.md           # Libraries & rationale
├─ frontend/                  # Vite + React app (served as static site)
└─ backend/                   # Go API (chi) with CORS, future jobs/connectors
```

**Principles**
- Keep the root clean. Source lives only in `frontend/` and `backend/`.
- Everything should build and run with a single `docker compose up` locally.
- Prefer composition over frameworks; add only what you need.

---

## 1) Frontend (Vite + React + TS)

### Folder Layout

```
frontend/
├─ Dockerfile
├─ nginx.conf
├─ index.html
├─ package.json
├─ tsconfig.json
├─ vite.config.ts
├─ .yarnrc.yml
├─ public/                    # (optional) static assets copied as-is
└─ src/
   ├─ app/                    # (optional) top-level app wiring (providers)
   │  └─ queryClient.ts       # TanStack Query client setup
   ├─ routes/                 # Tiny routers or file-organized route components
   │  └─ Home.tsx
   ├─ tiles/                  # Reusable "tile" UIs (CO2Tile, AQITile,...)
   │  ├─ CO2Tile.tsx
   │  └─ index.ts
   ├─ charts/                 # Vega-Lite specs & helpers
   │  ├─ co2.spec.ts          # exports a Vega-Lite spec for CO₂
   │  └─ helpers.ts           # axis/format helpers
   ├─ maps/                   # MapLibre layers, sources, styles
   ├─ components/             # Generic UI components (Buttons, Cards, Legend)
   ├─ hooks/                  # Custom React hooks (useTileData, useSSE,...)
   ├─ lib/                    # API client, formatters, date utils
   │  ├─ api.ts               # fetch/ky client; baseURL detection
   │  ├─ sse.ts               # SSE helper for live tiles
   │  └─ time.ts              # date-fns helpers
   ├─ schemas/                # zod schemas for API payloads
   ├─ styles/                 # (optional) tailwind.css, tokens, global.css
   ├─ types/                  # shared TS types (derived from zod when possible)
   ├─ __tests__/              # unit/component tests
   ├─ main.tsx                # entry (mounts <App/>)
   └─ App.tsx                 # root component
```

**Notes & Guardrails**
- **Tiles** are self-contained: each has UI + data hook + (optional) chart spec.
- **Charts**: prefer **Vega-Lite** specs (export functions returning specs given data). Put bespoke/D3 hero visuals in their own files.
- **API client**: `lib/api.ts` centralizes base URL resolution and error handling.
- **Validation**: parse API responses with **zod** in `schemas/` and use inference for TS types.
- **State**: use **TanStack Query** for async state; only introduce global state (Zustand/Valtio) if truly necessary.
- **Routing**: simple component routing is fine; you can add TanStack Router later without reshuffling.
- **Naming**: `PascalCase.tsx` for components; `kebab-case.ts` for utilities & specs.

**Example: CO₂ Tile minimal structure**
```
src/
  tiles/CO2Tile.tsx          # renders value + sparkline
  charts/co2.spec.ts         # vega-lite spec factory
  hooks/useCO2.ts            # fetches /api/metrics/co2 with React Query
  schemas/co2.ts             # zod schema for backend payload
```

---

## 2) Backend (Go, chi)

### Folder Layout (evolves from the current minimal main.go)

```
backend/
├─ Dockerfile
├─ go.mod
├─ main.go                   # entrypoint (can move to cmd/api/main.go later)
├─ .dockerignore
└─ internal/                 # (create as the app grows)
   ├─ http/                  # handlers, router, middlewares
   │  ├─ router.go
   │  └─ handlers/
   │     ├─ health.go
   │     ├─ metrics.go       # /api/metrics/{slug}
   │     └─ sse.go           # /stream/tiles (Server-Sent Events)
   ├─ connectors/            # NOAA, NSIDC, FIRMS, OpenAQ, etc.
   │  ├─ co2_noaa.go
   │  ├─ sea_ice_nsidc.go
   │  └─ ...
   ├─ jobs/                  # cron tasks (pollers); schedule + task funcs
   │  ├─ scheduler.go
   │  └─ tasks.go
   ├─ store/                 # DB/Cache abstraction (pgx/sqlc, redis)
   │  ├─ postgres.go
   │  ├─ redis.go
   │  └─ observations.sqlc.sql  # sqlc queries (if using sqlc)
   ├─ domain/                # core business types (Metric, Observation)
   ├─ config/                # env parsing, defaults
   └─ telemetry/             # OTEL / Prometheus setup
```

**Notes & Guardrails**
- Keep all importable code in `internal/` so it’s **not** imported by other modules.
- `connectors/` do **fetch + validate + normalize**. They do **not** write HTTP responses.
- `jobs/` orchestrate periodic pulls (e.g., robfig/cron). They **call** connectors then write to `store/`.
- `http/` depends on `store/` for reads; do **not** import `connectors/` directly from handlers.
- Prefer **append-only** `observations` storage; compute rollups on read (or materialize later).
- Expose clear endpoints:
  - `GET /health`
  - `GET /api/hello`
  - `GET /api/metrics/{slug}` → latest snapshot or small window
  - `GET /api/series/{slug}` → time series (with `?from=&to=`)
  - `GET /stream/tiles` → SSE for live tiles

**Conventions**
- **Metric slugs**: kebab-case, e.g., `co2-daily`, `temp-anomaly`, `gmsl`, `sea-ice-extent`, `aqi-city`, `wildfires`, `enso-oni`.
- **Validation**: sanitize upstream data; reject outliers using simple stats (z-score vs rolling window) before persisting.
- **Licensing**: each `connectors/*` file carries `source_url`, `license`, and `update_cadence` constants for traceability.

---

## 3) Cross-cutting Conventions

### Imports & Dependence
- **Frontend**: `tiles/` & `routes/` may import `components/`, `charts/`, `hooks/`, `lib/`, `schemas/`. Avoid circular deps.
- **Backend**: `http/` → `store/` & `domain/`; `jobs/` → `connectors/` + `store/`. **Never** `connectors/ → http/`.

### Error Handling
- Frontend: use `react-error-boundary` per tile; show “data unavailable” state w/ retry.
- Backend: return JSON errors `{ error: string }`, log with `slog` at the edge of the system.

### Testing
- Frontend: `__tests__/` colocated near features; Vitest + RTL.
- Backend: `*_test.go` colocated; use `httptest` for handlers and table-driven tests.

### Environment
- Public site URL: `VITE_PUBLIC_SITE_URL` baked at FE build.
- API PORT: `PORT` (default 8080).
- Future: add `.env` files and pass via compose or your cloud provider’s secrets.

---

## 4) Data Flow (high-level)

```
Upstream (NOAA/NSIDC/...) → connector fetch → validate/normalize →
append to storage (future: Postgres/Timescale + S3 snapshots) →
API endpoints → FE tiles (Query/SSE) → charts (Vega-Lite) & maps (MapLibre)
```

- For MVP, you can keep data in-memory or cache; graduate to Postgres when you add historical views and rollups.

---

## 5) Evolution Path

- **Stage 1 (MVP)**: `main.go` with handlers; no DB; live endpoints for CO₂, sea-ice, AQI, fires (reads pass-through from upstream with short cache).
- **Stage 2**: add `internal/` tree; introduce `jobs/` + `connectors/` and a `db` service with Postgres/Timescale.
- **Stage 3**: public `/api/series/*`, embeddable tiles, and a small admin to toggle refresh intervals.
- **Stage 4**: observability (OTEL/Prometheus), strict CORS, rate limiting, and auth for write/admin endpoints.

---

## 6) Quick Scaffolding Commands

Create the suggested FE folders:
```bash
mkdir -p frontend/src/{{app,routes,tiles,charts,maps,components,hooks,lib,schemas,styles,__tests__}}
```

Create the suggested BE folders (when you graduate from single-file main):
```bash
mkdir -p backend/internal/{{http/handlers,connectors,jobs,store,domain,config,telemetry}}
```

---

## 7) Checklist Before Shipping
- [ ] Each tile shows **source, cadence, last updated**
- [ ] CORS restricts to `https://one-earth.info`
- [ ] 4xx/5xx JSON error shapes standardized
- [ ] Rate limit `/api/*`
- [ ] Cache headers set (5–30 min) for public GETs
- [ ] README has **run**, **deploy**, and **licensing** notes
