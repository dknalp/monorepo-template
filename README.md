# Agency Monorepo Template

Production-ready monorepo starter for Go + Next.js + Tailwind CSS projects.

## Stack

| Layer | Technology |
|---|---|
| Frontend | Next.js 15, React 19, Tailwind CSS 4 |
| Backend | Go 1.22, stdlib net/http |
| Monorepo (JS) | Turborepo + pnpm workspaces |
| Monorepo (Go) | go.work (native Go workspaces) |
| Database | PostgreSQL (one DB per service) |
| Containers | Docker + Docker Compose |

## Structure

```
.
├── apps/
│   ├── web/          # Main Next.js app (port 3000)
│   └── admin/        # Admin Next.js app (port 3001)
├── services/
│   ├── api/          # API gateway service (port 8080)
│   ├── auth/         # Auth service (port 8081)
│   └── storage/      # Storage service (port 8082)
├── packages/
│   ├── go-api-lib/   # Shared Go middleware, response helpers, errors
│   └── ui/           # Shared React components + Tailwind preset
├── infra/
│   ├── docker-compose.yml
│   └── dockerfiles/
├── go.work           # Go workspace (ties all Go modules)
├── turbo.json        # Turborepo config
└── pnpm-workspace.yaml
```

## Rules for Every Project

1. **`cmd/<service>/main.go` is bootstrap only** — no business logic
2. **Database-per-service** — never share schemas across services
3. **Shared Go code** goes in `packages/go-api-lib`
4. **Shared UI code** goes in `packages/ui`
5. **`internal/` is private** — Go enforces this at compile time

## Getting Started

```bash
# Install JS dependencies
pnpm install

# Start all frontends
pnpm dev

# Start a Go service
cd services/api && go run ./cmd/api

# Start everything with Docker
cd infra && docker compose up
```

## Adding a New Service

```bash
mkdir -p services/myservice/cmd/myservice
mkdir -p services/myservice/internal/{handler,service,repository}
# Copy go.mod from services/api, update module name
# Add to go.work
```

## Adding a New Frontend App

```bash
cd apps && pnpm create next-app myapp
# Add @agency/ui as dependency
# Extend tailwind-preset from packages/ui
```
