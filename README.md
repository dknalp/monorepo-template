# Agency Monorepo Template

> Production-ready monorepo starter for Go + Next.js + Tailwind CSS.
> Clone this for every new project. Read `CONTRIBUTING.md` before writing any code.

[![CI](https://github.com/dknalp/monorepo-template/actions/workflows/ci.yml/badge.svg)](https://github.com/dknalp/monorepo-template/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

---

## Stack

| Layer | Technology |
|---|---|
| Frontend | Next.js 15, React 19, Tailwind CSS 4 |
| Backend | Go 1.22, stdlib `net/http` |
| Monorepo (JS) | Turborepo + pnpm workspaces |
| Monorepo (Go) | `go.work` (native Go workspaces) |
| Database | PostgreSQL — **one DB per service** |
| Containers | Docker + Docker Compose |
| CI | GitHub Actions |

---

## Structure

```
.
├── apps/
│   ├── web/               # Main Next.js app          → localhost:3000
│   └── admin/             # Admin Next.js app          → localhost:3001
│
├── services/
│   ├── api/               # API gateway               → localhost:8080
│   ├── auth/              # Auth service              → localhost:8081
│   └── storage/           # Storage service           → localhost:8082
│
├── packages/
│   ├── go-api-lib/        # Shared Go: middleware, response helpers, errors
│   └── ui/                # Shared React components + Tailwind preset
│
├── infra/
│   ├── docker-compose.yml
│   └── dockerfiles/       # One Dockerfile per service
│
├── .github/workflows/ci.yml
├── go.work                # Ties all Go modules together
├── turbo.json
└── pnpm-workspace.yaml
```

---

## Getting Started

### Prerequisites

- Go 1.22+
- Node.js 20+
- pnpm 9+
- Docker + Docker Compose

### Local Development

```bash
# Clone
git clone https://github.com/dknalp/monorepo-template.git my-project
cd my-project

# Install JS deps
pnpm install

# Run all frontends (hot reload)
pnpm dev

# Run a Go service
cd services/api && go run ./cmd/api

# Run everything with Docker
cd infra && docker compose up
```

### First-time Setup for a New Project

```bash
# 1. Update module names in all go.mod files
#    Change: github.com/agency/* → github.com/yourorg/yourproject/*

# 2. Update package names in package.json files
#    Change: @agency/* → @yourorg/*

# 3. Copy .env.example → .env and fill in values

# 4. Run go work sync
go work sync
```

---

## The 5 Rules (must be followed on every project)

1. **`cmd/<service>/main.go` is bootstrap only** — no business logic ever
2. **Database-per-service** — never share schemas across services
3. **Shared Go code** → `packages/go-api-lib` only
4. **Shared UI code** → `packages/ui` only
5. **`internal/` is private** — Go enforces this at compile time

See [CONTRIBUTING.md](./CONTRIBUTING.md) for full standards.

---

## Adding a New Service

```bash
mkdir -p services/myservice/cmd/myservice
mkdir -p services/myservice/internal/{handler,service,repository}
# Copy go.mod from services/api, update module name
# Add entry to go.work
# Add service to infra/docker-compose.yml
# Add infra/dockerfiles/myservice.Dockerfile
```

## Adding a New Frontend App

```bash
cd apps && pnpm create next-app myapp --typescript --tailwind --app --no-src-dir
# Add "@agency/ui": "workspace:*" to dependencies
# Replace tailwind.config.ts to extend packages/ui/tailwind-preset
```

---

## License

[MIT](./LICENSE)
