# Contributing Guide

This document defines the standards every worker must follow on all agency projects.

## Golden Rules

1. **`cmd/<service>/main.go` is bootstrap only** — wire deps, start server. No business logic.
2. **Database-per-service** — each service owns its own DB. Never query another service's DB directly.
3. **Cross-service communication** — HTTP or gRPC only. No shared DB schemas, no shared memory.
4. **Shared Go code** → `packages/go-api-lib`. Shared UI code → `packages/ui`. Never copy-paste between services.
5. **`internal/` is private** — Go enforces this. Anything that must be shared goes in `packages/`.

## Branch Naming

```
feat/short-description
fix/short-description
chore/short-description
```

## Commit Format

```
feat: add user profile endpoint
fix: correct token expiry calculation
chore: update Go to 1.22.3
```

## Adding a New Go Service

```bash
# 1. Create the directory structure
mkdir -p services/myservice/cmd/myservice
mkdir -p services/myservice/internal/{handler,service,repository}

# 2. Create go.mod (copy from services/api, change module name)
# module github.com/agency/myservice

# 3. Add to go.work
#    use ./services/myservice

# 4. Add to infra/docker-compose.yml
# 5. Add a Dockerfile under infra/dockerfiles/myservice.Dockerfile
```

## Adding a New Frontend App

```bash
cd apps
pnpm create next-app myapp --typescript --tailwind --app --no-src-dir

# Add to apps/myapp/package.json dependencies:
# "@agency/ui": "workspace:*"

# Replace tailwind.config.ts to extend the shared preset from packages/ui
```

## Adding Shared Go Code

Add to `packages/go-api-lib/`. All services get it automatically via `go.work`.
Never add service-specific logic here — only truly shared utilities.

## Adding Shared UI Components

Add to `packages/ui/components/`. Export from `packages/ui/index.ts`.
All Next.js apps get it automatically via pnpm workspace.

## Environment Variables

- Never hardcode ports, DB URLs, or secrets.
- All env vars must be documented in `.env.example`.
- Services fall back to sensible dev defaults when env vars are missing.

## CI

All PRs must pass:
- `go vet` + `go build` for all services
- `go test ./...`
- `pnpm type-check` + `pnpm lint`
- Docker build check for all service images
