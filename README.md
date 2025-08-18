# One Earth — Vite + React (Yarn) + Go (Docker)

Domain: **https://one-earth.info**

This is a slim Docker-first starter: static **Vite + React** frontend (served by nginx) and a tiny **Go** API (chi) with CORS enabled.

## Prereqs
- Docker Desktop or Docker Engine (Compose v2)

## Quick start
```bash
cd one-earth
docker compose build --no-cache
docker compose up
```

Open:
- Frontend → http://localhost:8080
- API health → http://localhost:8081/health
- API hello  → http://localhost:8081/api/hello
