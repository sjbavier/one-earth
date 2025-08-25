# One Earth — Project Summary

_Last updated: {{today’s date}}_

## Overview

**One Earth** is a data visualization website focused on **climate change and environmental issues**, inspired by the visual clarity of Visual Capitalist but grounded in **live, reputable data**.

Domain: [one-earth.info](https://one-earth.info)  
GitHub repo: `git@github.com:sjbavier/one-earth.git`

---

## Tech Stack

(see [`tech-stack.md`](./tech-stack.md) for details)

- **Frontend:** React 18 + Vite, Yarn 4 (Corepack), TanStack Query, Vega-Lite via `react-vega`, Tailwind-like tokens, shadcn/ui (optional).
- **Backend:** Go (Chi, CORS), REST APIs fetching authoritative CSV/JSON (NOAA, NASA, NSIDC, etc.).
- **Infra:** Docker Compose (frontend, backend, nginx), deployable to AWS/GCP/Fly.
- **Data:** Direct from reputable public sources, cached server-side (TTL), safe licensing (NOAA/NASA public domain, Copernicus CC-BY).

---

## Directory Structure

(see [`directory-structure.md`](./directory-structure.md))

- `frontend/` → React + Vite app (tiles, hooks, charts, themes).
- `backend/` → Go API (connectors, handlers, caching).
- `docs/` → All architecture and planning docs.

---

## Data Visualizations

(see [`viz-ideas.md`](./viz-ideas.md))

**MVP metrics (live tiles first):**

1. Atmospheric CO₂ (NOAA GML)
2. Global temperature anomaly (NASA GISTEMP)
3. Global mean sea level (NASA/AVISO)
4. Sea-ice extent (NSIDC)
5. Air quality (OpenAQ/AirNow)
6. Active wildfires (NASA FIRMS)
7. Grid carbon intensity (Electricity Maps, EIA, UK CI)
8. Deforestation alerts (Global Forest Watch)
9. ENSO indices (NOAA CPC)
10. Annual emissions (Global Carbon Budget / OWID)

Each tile: value + sparkline → “last updated” → provenance (source + license).

---

## Current Focus

(see [`todo-co2.md`](./todo-co2.md))

**Atmospheric CO₂ (NOAA GML)**

- ✅ Define backend connector & endpoints.
- ✅ Define frontend schemas, hooks, tile plan.
- 🚧 Implement sparkline tile in frontend.
- 🚧 Integrate into App grid.

---

## Theme

(see `frontend/src/styles/theme.css` and `frontend/src/charts/theme.ts`)

- **Light mode:** Reuters-style (paper, muted, teal/rust accents).
- **Dark mode:** Apple HIG-style (iOS dark, system colors).
- Toggle via `ThemeToggle` component.

---

## Workflow

- Yarn 4 via Corepack (`packageManager: "yarn@4.3.1"`).
- Commit `yarn.lock` for consistent dependencies.
- Docker Compose builds frontend/backend for local dev.
- PRs grouped by feature slice (backend, frontend, integration).

---

## Next Steps

1. Finish **CO₂ tile** (backend endpoints + frontend tile).
2. Add **temperature anomaly** tile.
3. Add more live tiles from MVP list.
4. Expand docs for seasonal cycle & spiral vizzes.
5. Deploy staging instance (AWS/GCP/Fly).

---

## References

- [`requirements.md`](./requirements.md)
- [`tech-stack.md`](./tech-stack.md)
- [`directory-structure.md`](./directory-structure.md)
- [`viz-ideas.md`](./viz-ideas.md)
- [`todo-co2.md`](./todo-co2.md)
