# EstateHub API

EstateHub API is a professional REST API foundation written in Go. The project is intentionally framework-free and uses `net/http` from the Go standard library for HTTP routing and middleware.

The repository is currently in **Phase 0: Professional Project Setup**. This phase prepares the application foundation before any business endpoints or CRUD features are implemented.

## Current Phase

**Phase 0: Professional Project Setup**

The goal of this phase is to establish a clean, production-oriented baseline for the API:

- Go project structure
- PostgreSQL dependency
- Docker Compose setup
- SQL migrations
- Environment-based configuration
- Structured logging with `log/slog`
- HTTP request logging middleware
- Health check endpoint
- Readiness check endpoint
- Prepared application layers
- Graceful shutdown
- No business endpoints yet

## Tech Stack

- Go `1.26.1`
- Go standard library HTTP server with `net/http`
- PostgreSQL
- Docker Compose
- `database/sql`
- `github.com/lib/pq` PostgreSQL driver
- `log/slog` for structured logging

## Project Structure

Current repository structure:

```text
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── platform/
│       ├── config/
│       │   └── config.go
│       ├── database/
│       │   ├── database.go
│       │   └── migations.go
│       ├── http/
│       │   ├── health.go
│       │   ├── middleware.go
│       │   ├── response.go
│       │   └── route.go
│       └── logger/
│           └── logger.go
├── migrations/
│   └── 001_init.sql
├── docker-compose.yml
├── go.mod
└── go.sum
```

### Main Components

- `cmd/api/main.go`: application entrypoint. It loads configuration, opens the database connection, runs migrations, builds the HTTP router, starts the server, and handles graceful shutdown.
- `internal/platform/config`: loads application settings from environment variables.
- `internal/platform/database`: opens the PostgreSQL connection and runs SQL migrations.
- `internal/platform/logger`: creates a structured `slog` logger.
- `internal/platform/http`: contains routing, health/readiness handlers, JSON responses, and request logging middleware.
- `migrations`: contains SQL migration files executed on application startup.

## Environment Variables

The application reads configuration from environment variables.

| Variable | Required | Default | Description |
|---|---:|---|---|
| `APP_NAME` | No | `EstateHub API` | Application name used in startup logs. |
| `APP_ENV` | No | `development` | Runtime environment. Production uses JSON logs; development uses text logs. |
| `APP_PORT` | No | `8080` | HTTP server port. |
| `DATABASE_URL` | Yes | none | PostgreSQL connection string. |
| `READINESS_TIMEOUT_SECONDS` | No | `2` | Timeout used by the readiness database ping. |

Example local environment:

```env
APP_NAME=EstateHub API
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgres://postgres:password@localhost:5432/estatehub?sslmode=disable
READINESS_TIMEOUT_SECONDS=2
```

## Running PostgreSQL With Docker Compose

The project includes a `docker-compose.yml` file intended to run PostgreSQL locally.

```bash
docker compose up -d postgres
```

Check container status:

```bash
docker compose ps
```

Stop the database:

```bash
docker compose down
```

### Current Docker Compose Status

The current `docker-compose.yml` needs a small fix before it can be considered complete:

- It references a `postgres_socket` volume that is not declared.
- The PostgreSQL healthcheck uses an invalid command and mismatched username.
- The container name has a typo: `estaehub-postgres`.

Until those are fixed, `docker compose config` reports the compose project as invalid.

## Running the API Locally

Set the required environment variables first. On Windows PowerShell:

```powershell
$env:APP_NAME="EstateHub API"
$env:APP_ENV="development"
$env:APP_PORT="8080"
$env:DATABASE_URL="postgres://postgres:password@localhost:5432/estatehub?sslmode=disable"
$env:READINESS_TIMEOUT_SECONDS="2"
```

Run the API:

```bash
go run ./cmd/api
```

On startup, the application:

1. Loads configuration from environment variables.
2. Opens a PostgreSQL connection.
3. Runs pending migrations from the `migrations` directory.
4. Creates the HTTP router.
5. Starts the HTTP server.
6. Waits for `SIGINT` or `SIGTERM` to gracefully shut down.

## Health Check

The intended Phase 0 health endpoint is:

```http
GET /healthz
```

Expected successful response:

```json
{
  "status": "ok"
}
```

Test command:

```bash
curl http://localhost:8080/healthz
```

### Current Route Status

The current code registers `/healtz` instead of `/healthz`. This should be fixed in `internal/platform/http/route.go` and `internal/platform/http/health.go`.

## Readiness Check

The intended Phase 0 readiness endpoint is:

```http
GET /readyz
```

Expected successful response when PostgreSQL is reachable:

```json
{
  "status": "ready",
  "checks": {
    "database": "available"
  }
}
```

Expected response when PostgreSQL is unavailable:

```json
{
  "status": "not_ready",
  "checks": {
    "database": "unavailable"
  }
}
```

Test command:

```bash
curl http://localhost:8080/readyz
```

### Current Route Status

The current code registers `/readz` instead of `/readyz`. This should be fixed before Phase 0 is considered complete.

## Health vs Readiness

The health check and readiness check have different responsibilities.

`/healthz` verifies that the API process is running and can respond to HTTP requests. It should stay lightweight and should not depend on external services.

`/readyz` verifies that the API is ready to serve traffic. In this project, readiness includes a PostgreSQL ping using `PingContext` with a configurable timeout.

This separation allows the process to be alive while still reporting that it is temporarily not ready because a dependency is unavailable.

## Migrations

Migrations are stored in the `migrations` directory as `.sql` files.

On application startup, `database.RunMigration`:

1. Ensures the `schema_migrations` table exists.
2. Finds all `.sql` files in the migrations directory.
3. Sorts migration files by filename.
4. Skips migrations already recorded in `schema_migrations`.
5. Runs each pending migration inside a transaction.
6. Records the applied migration version.

The current migration file is:

```text
migrations/001_init.sql
```

It creates:

- `schema_migrations`
- `orders`
- `customers`
- `products`

These tables prepare the database for later phases. No business HTTP endpoints exist yet.

## Structured Logging

Logging is implemented with Go's `log/slog` package.

The logger configuration is environment-aware:

- `APP_ENV=production`: JSON logs at info level.
- Other environments: text logs at debug level.

The HTTP middleware logs request details, including:

- HTTP method
- Request path
- Response status
- Duration in milliseconds
- Remote address

Example log fields:

```text
method=GET path=/readyz status=200 duration_ms=3 remote_addr=127.0.0.1:50000
```

## Application Architecture

The application is organized around a simple layered structure:

- `cmd/api`: composition root and process lifecycle.
- `internal/platform/config`: configuration loading.
- `internal/platform/database`: database connection and migrations.
- `internal/platform/logger`: structured logger setup.
- `internal/platform/http`: HTTP transport, routes, middleware, and platform endpoints.
- `migrations`: database schema changes.

The intended next application modules are:

- `internal/order`
- `internal/customer`
- `internal/product`

These modules are not present yet in the current repository.

## Graceful Shutdown

The API listens for `SIGINT` and `SIGTERM`.

When a shutdown signal is received, the server calls `Shutdown` with a timeout context. This allows in-flight requests to finish before the process exits.

The current shutdown timeout is:

```text
10 seconds
```

## Current Limitations

- No business endpoints exist yet.
- No CRUD handlers exist yet.
- No request validation layer exists yet.
- No tests are currently present.
- `.env.example` is missing.
- `Makefile` is not present.
- `internal/order`, `internal/customer`, and `internal/product` are not present yet.
- The current route names are misspelled as `/healtz` and `/readz`.
- The readiness handler should return immediately after writing a `503` response.
- The request logging middleware should use `WriteHeader`, not `writeHeader`, to capture non-200 status codes correctly.
- The request logging middleware currently accepts a logger but calls package-level `slog.Info`.
- `internal/platform/database/migations.go` should be renamed to `migrations.go`.
- `docker-compose.yml` currently needs fixes before Docker Compose can validate it.

## Pending Items

Before Phase 0 can be considered complete:

- Add `.env.example`.
- Fix Docker Compose volume and healthcheck configuration.
- Rename `/healtz` to `/healthz`.
- Rename `/readz` to `/readyz`.
- Fix readiness handler control flow after database ping failure.
- Fix request logging middleware status capture.
- Use the injected `*slog.Logger` in request logging middleware.
- Rename `migations.go` to `migrations.go`.
- Add prepared module directories for:
  - `internal/order`
  - `internal/customer`
  - `internal/product`

## Next Phase Roadmap

Phase 1 should start only after Phase 0 is complete.

Planned next steps:

- Define the first business module boundaries.
- Add domain models for `order`, `customer`, and `product`.
- Add repository interfaces where needed.
- Add service/application layer behavior.
- Add HTTP handlers for business use cases.
- Add request/response DTOs.
- Add tests for handlers, services, and database behavior.
- Keep routing based on `net/http`.

## Useful Commands

Run tests:

```bash
go test ./...
```

Run the API:

```bash
go run ./cmd/api
```

Start PostgreSQL:

```bash
docker compose up -d postgres
```

Stop PostgreSQL:

```bash
docker compose down
```

Validate Docker Compose:

```bash
docker compose config
```

Test the intended health endpoint:

```bash
curl http://localhost:8080/healthz
```

Test the intended readiness endpoint:

```bash
curl http://localhost:8080/readyz
```

Format Go code:

```bash
gofmt -w .
```

