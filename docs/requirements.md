# One Earth — Requirements

**Domain:** https://one-earth.info  
**Stack:** Vite + React (Yarn) frontend served by nginx, Go (chi) API, Docker-first, no SSR for MVP.

---

## 1) Vision & Goals

Build a beautiful, credible, and near‑real‑time **climate data visualization** site. Show live tiles for key indicators, each with transparent sourcing, licensing, and “last updated” timestamps. Prioritize clarity, speed, and trust.

**Success signals (MVP):**

- Home loads fast (LCP < 2.5s on mid‑tier mobile).
- 4–6 live tiles auto-refresh with stable cadence and visible provenance.
- Each tile has **Download data**, **Source**, and **Method** links.
- API stable under light load with edge caching and rate limits.

---

## 2) MVP Scope

### 2.1 Live tiles (initial)

- **CO₂ (ppm)** — NOAA GML (daily prelim → weekly/monthly official)
- **AQI (city PM2.5 / AQI)** — OpenAQ (global) + AirNow (US, API key)
- **Wildfires (active points)** — NASA FIRMS (VIIRS/MODIS, NRT/URT)
- **Sea‑ice extent** — NSIDC Sea Ice Index (daily)
- (Stretch) **Global temp anomaly** — NASA GISTEMP v4 (monthly)
- (Stretch) **Global mean sea level** — NASA Sea Level (monthly)

> Data source details live in `docs/sources.md`.

### 2.2 Pages

- **Home (Dashboard):** responsive grid of tiles, each links to a detail page.
- **Tile Detail:** small explainer, mini time series, source & license.
- **About the data:** methodology pages per metric.

### 2.3 Out of scope (MVP)

- SSR/RSC, auth, DB persistence (optional later), CMS, complex user settings.

---

## 3) Functional Requirements

### 3.1 Frontend

- **Tiles** show: title, value, unit, trend (sparkline/arrow), **Last updated**, **Source**, **License**, **Download**.
- **Updates**: tiles poll or subscribe (SSE) without full page reload.
- **Accessibility**: keyboard navigable, contrast-safe palettes, labeled charts.
- **Shareability**: basic meta tags for title/description and Open Graph image (static placeholder ok for MVP).

### 3.2 API

- **Endpoints (MVP):**
  - `GET /health` → `"ok"`
  - `GET /api/hello` → greeting JSON
  - `GET /api/metrics/:slug` → latest snapshot
  - `GET /api/series/:slug?from=&to=&limit=&granularity=` → small time window (MVP may proxy upstream or return a 7–30 day sample)
  - `GET /stream/tiles` (optional MVP) → SSE for push updates
- **Responses** include: `value`, `unit`, `updated_at`, `source_name`, `source_url`, `license`, `notes` (when relevant).
- **Caching**: set `Cache-Control` (5–30 min), pass through `ETag`/`Last-Modified` when possible.
- **Rate limiting**: per IP token bucket (e.g., 60 req/min) for public routes.
- **Errors**: JSON `{ "error": "message" }` with appropriate status codes.

**OpenAPI sketch (excerpt):**

```yaml
openapi: 3.0.3
info:
  title: One Earth API
  version: 0.1.0
paths:
  /health:
    get:
      summary: Health check endpoint
      responses:
        "200":
          description: OK
  /api/metrics/{slug}:
    get:
      summary: Get latest metric snapshot by slug
      parameters:
        - in: path
          name: slug
          schema:
            type: string
          required: true
          description: Metric slug identifier
      responses:
        "200":
          description: Latest metric snapshot
          content:
            application/json:
              schema:
                type: object
                properties:
                  slug:
                    type: string
                    example: co2-daily
                  value:
                    type: number
                    example: 420.5
                  unit:
                    type: string
                    example: ppm
                  updated_at:
                    type: string
                    format: date-time
                    example: "2025-08-29T00:00:00Z"
                  source_name:
                    type: string
                    example: NOAA GML
                  source_url:
                    type: string
                    example: https://gml.noaa.gov/ccgg/trends/
                  license:
                    type: string
                    example: NOAA public domain
        "404":
          description: Metric not found
```

### 3.3 Data fetchers (MVP; simple mode)

- Fetch live from upstream on request with short caching (edge/CDN and in‑process).
- Validate and **sanitize** responses. Reject obviously broken/outlier values.
- Log fetch time, status, and headers (`ETag`, `Last-Modified`).
- **Optional**: in‑memory ring buffer per metric for the last N points to draw sparklines.

---

## 4) Non‑Functional Requirements

### Performance

- **Frontend**: initial JS < 180KB gz (MVP), LCP < 2.5s, CLS < 0.1, TBT < 200ms.
- **API**: p95 latency < 200ms for cached hits; < 600ms for cold upstream calls.

### Reliability

- Uptime target MVP 99.5%. Graceful degradation: show “data currently unavailable” instead of blank.

### Security

- CORS allowlist: `https://one-earth.info` (dev: `http://localhost:8080`).
- Set security headers (CSP, HSTS, X-Content-Type-Options, Referrer-Policy).
- Do not expose secrets in frontend; keep API keys server-side.

### Accessibility

- WCAG 2.1 AA color/contrast, focus states, keyboard navigation, ARIA labels for charts.
- Color‑blind‑safe palettes (no red vs green only).

### Observability

- Structured logs (Go `slog`).
- Error reporting (Sentry) FE/BE.
- Metrics: Prometheus `/metrics`; traces via OpenTelemetry (later).

### SEO/Meta

- Static meta tags per page; pre-rendered HTML from nginx is fine (no SSR required).
- OG/Twitter image: static placeholder ok for MVP.

### Privacy & Analytics

- Privacy-first analytics (e.g., Plausible/Umami), with IP anonymization and cookieless mode.

---

## 5) Architecture Overview

**Frontend** (Vite + React):

- Fetch via `fetch`/`ky` + **TanStack Query**; validate with **zod**.
- Visuals: **Vega‑Lite** (`react-vega`) and **MapLibre GL** for maps.
- Optional SSE client helper for truly live tiles.

**API** (Go, chi):

- Handlers expose metrics endpoints.
- Simple fetcher layer per source with retry/backoff and basic validation.
- Caching headers + optional in‑memory cache.
- (Later) background jobs + Postgres/Timescale + Redis for persistence and queues.

**CDN/Edge**: cache `/api/*` GETs for 5–30 min; purge on new data if background jobs are added.

---

## 6) Directory Structure (contract)

See `docs/directory-structure.md` for full details. Short form:

```
frontend/
  src/{tiles,charts,maps,components,hooks,lib,schemas,styles}
backend/
  main.go  # (Stage 1)
  internal/{http,connectors,jobs,store,domain,config,telemetry}  # (Stage 2)
docs/{sources.md,tech-stack.md,directory-structure.md,requirements.md}
```

---

## 7) Environments & Config

**Local (Docker compose)**

- FE: http://localhost:8080
- API: http://localhost:8081

**Env vars**

- Frontend build: `VITE_PUBLIC_SITE_URL` (baked at build)
- API: `PORT` (default 8080)
- (Later) `ALLOWED_ORIGINS`, `RATE_LIMIT`, API keys for upstream providers

**Deploy targets (suggested)**

- FE: Cloudflare Pages / Netlify (static)
- API: Fly.io / Cloud Run (Docker image)
- (Later) DB: Timescale Cloud / Neon; Object storage: R2/S3

---

## 8) Testing Strategy

**Frontend**

- Unit: components & utils (Vitest, RTL).
- Integration: tile hooks with mocked fetch (msw).
- E2E (later): Playwright for key flows (load dashboard, tile refresh, source link).

**Backend**

- Unit: handler tests with `httptest`.
- Integration: upstream fetchers hitting mock servers (recorded fixtures).
- Contract: zod schemas in FE aligned with OpenAPI (generate TS types later).

**Acceptance checks (automatable)**

- `/health` returns 200.
- `/api/metrics/co2` returns JSON with `value`, `unit`, `updated_at`, `source_name`.
- Dashboard displays value and “Last updated” within expected cadence.
- Tiles render accessible labels; keyboard focus order is correct.

---

## 9) MVP Acceptance Criteria (per tile)

**Common**

- [ ] Shows current value + unit.
- [ ] Displays **Last updated** (UTC) and **Source** link.
- [ ] “Download” links the exact CSV/JSON used.
- [ ] Handles error state within 2s and shows retry.
- [ ] Updates on a sensible cadence (CO₂ daily; Sea‑ice daily; AQI hourly; FIRMS 5–15 min).

**CO₂**

- [ ] NOAA CSV parsed; value matches upstream within tolerance.
- [ ] Trend sparkline shows last 30–60 points.

**AQI**

- [ ] City selector; shows PM2.5 and AQI (or PM2.5 if AQI not available).
- [ ] Falls back gracefully if a station is offline.

**Wildfires**

- [ ] Map shows active points; clicking reveals age/intensity attributes.
- [ ] Auto-refresh (5–15 min) without page reload.

**Sea‑ice**

- [ ] Daily extent chart (last 30–90 days) with anomaly vs. median band (if available).

---

## 10) Backlog (Post‑MVP)

- Add **Deforestation alerts** (GFW RADD/GLAD).
- Add **ENSO** tile (ONI, Niño 3.4).
- Grid carbon intensity (UK free API; Electricity Maps if budget).
- Persist observations in Postgres/Timescale; publish `/api/series/*`.
- Public embeddable tiles (iframe or JS snippet).
- Admin dashboard for data freshness and job controls.
- OpenAPI → generate TS client; add typed SDK for FE.

---

## 11) Risks & Mitigations

- **Upstream instability** → use retries/backoff + short cache + graceful UI fallbacks.
- **API key limits** (AirNow, others) → cache at edge; consider server-side aggregator.
- **Payload drift** → validate with zod; monitor with Sentry/alerts.
- **Performance regressions** → budgets + CI checks; track Web Vitals.

---

## 12) Runbook (local)

```bash
# build & run all
docker compose build --no-cache
docker compose up

# endpoints
curl -s http://localhost:8081/health
curl -s http://localhost:8081/api/hello
```

---

_Last updated: 2025-08-19_
