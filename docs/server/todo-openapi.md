# TODO: Add OpenAPI Spec and Swagger UI Integration

This document outlines the detailed tasks to add a complete OpenAPI specification file, integrate Swagger UI into the React frontend, and serve the OpenAPI spec file either as a static asset or via an API endpoint.

---

## 1. Create OpenAPI Specification File

- [ ] Define a complete OpenAPI 3.0+ spec file (`openapi.yaml` or `openapi.json`) describing all backend API endpoints:
  - `/health`
  - `/api/hello`
  - `/api/metrics/{slug}`
  - `/api/series/{slug}`
  - `/stream/tiles` (SSE)
- [ ] Include request parameters, response schemas, status codes, and example payloads.
- [ ] Validate the spec file using an OpenAPI validator tool.
- [ ] Place the spec file in a suitable location:
  - Option A: `frontend/public/` as a static asset
  - Option B: `backend/internal/http/assets/` if serving via backend API

---

## 2. Serve OpenAPI Spec File

- [ ] If serving from backend:
  - Add a new HTTP handler in Go to serve the OpenAPI spec file at e.g. `/api/openapi.yaml`.
  - Ensure proper content-type headers (`application/yaml` or `application/json`).
- [ ] If serving from frontend:
  - No backend changes needed; the spec file will be accessible as a static asset.

---

## 3. Integrate Swagger UI in React Frontend

- [ ] Add `swagger-ui-react` package to frontend dependencies.
- [ ] Create a new React component (e.g., `SwaggerUI.tsx`) that:
  - Imports and renders Swagger UI React component.
  - Loads the OpenAPI spec file from the static asset or backend endpoint.
- [ ] Add a new route or page in the frontend app to display the Swagger UI component.
- [ ] Style the Swagger UI to match the existing frontend theme if desired.

---

## 4. Testing and Validation

- [ ] Verify the OpenAPI spec loads correctly in Swagger UI.
- [ ] Test all documented endpoints via the Swagger UI interface.
- [ ] Ensure CORS and security headers allow fetching the spec file.
- [ ] Confirm the integration works in local development and production builds.

---

## 5. Documentation and Cleanup

- [ ] Update project documentation to mention the new API docs page.
- [ ] Add instructions for maintaining and updating the OpenAPI spec.
- [ ] Remove any temporary or unused files related to the integration.

---

## 6. Known Issues and Research

- [ ] Investigate and resolve TypeScript errors with `swagger-ui-react` component usage.
- [ ] Address runtime error: `Cannot read properties of undefined (reading 'from')` in `safe-buffer` within `swagger-ui-react`.
- [ ] Explore alternative integration approaches such as using `swagger-ui-dist` with a React wrapper.
- [ ] Research dependency conflicts and bundler compatibility with Vite and React 18.
- [ ] Document findings and recommended solutions for Swagger UI integration issues.

## Optional Enhancements

- [ ] Consider adding Redoc as an alternative or complement to Swagger UI.
- [ ] Automate OpenAPI spec generation from Go code using tools like `swaggo` or `go-swagger`.
- [ ] Add versioning to the OpenAPI spec and docs route.

---

This checklist provides a clear roadmap to implement interactive API documentation for the One Earth service using OpenAPI and Swagger UI.
