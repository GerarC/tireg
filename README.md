# tireg

A backend to use with models, used for maintain a local database for time registration.

## Architecture

The project follows a hexagonal (ports & adapters) architecture organized by
vertical business modules under `internal/` (`health`, `user`, `auth`, and
more to come: `project`, `glossary`, `time-registry`), plus a `shared`
module for cross-cutting code. Full conventions and rules are documented in
[CLAUDE.md](CLAUDE.md).

Relevant architectural decisions are recorded as ADRs in
[documentation/adr/](documentation/adr/):

- [0001 - Hexagonal architecture organized by business module](documentation/adr/0001-hexagonal-architecture-by-module.md)
- [0002 - Custom reflection-based dependency injection container](documentation/adr/0002-custom-reflection-based-di-container.md)
- [0003 - Postgres access via pgx and a plain SQL schema](documentation/adr/0003-postgres-access-via-pgx-and-plain-sql-schema.md)
- [0004 - JWT authentication consuming the user module's api](documentation/adr/0004-jwt-authentication-consuming-user-api.md)
- [0005 - Unified domain error structure](documentation/adr/0005-unified-domain-error-structure.md)

## Requirements

- Go 1.25.4+
- Docker + Docker Compose (for running with Postgres)

> **Windows:** run everything (build, run, test) inside **WSL**, not
> directly in PowerShell/cmd. Windows Smart App Control blocks locally
> built, unsigned `.exe` files (including `go run` and `go test` binaries),
> which makes native Windows execution unreliable. WSL binaries are plain
> Linux ELF executables and are unaffected.

## Running with Docker (recommended)

```bash
docker compose up --build
```

This starts the app on `:8080` and a Postgres 16 instance, with the schema
in `db/init/` applied automatically on first startup.

## Running the server locally (without Docker)

Requires a reachable Postgres instance; configure it via `.env` (copy
`.env.example`) or environment variables.

```bash
go build -o tireg ./cmd/tireg && ./tireg
```

The server listens on `:8080` by default (override with the `PORT` env var).
A local `.env` file (see `.env.example`) is also loaded on startup, without
overriding real environment variables.

## API

Interactive Swagger UI: `http://localhost:8080/swagger/index.html`.

```bash
curl http://localhost:8080/api/v1/health

curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"first_name":"Ada","last_name":"Lovelace","username":"ada","email":"ada@example.com","password":"super-secret"}'

curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"identifier":"ada","password":"super-secret"}'
```

Errors share one JSON shape across the whole API (see
[ADR 0005](documentation/adr/0005-unified-domain-error-structure.md)):

```json
{
  "code": 409,
  "message": "USER_ALREADY_TAKEN",
  "details": ["username already taken", "email already registered"],
  "timestamp": "2026-09-01T12:00:00Z"
}
```

## Testing

Tests live under `test/`, mirroring the structure of `internal/`:

```bash
go test ./...
```
