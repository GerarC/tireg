# tireg

A backend to use with models, used for maintain a local database for time registration.

## Architecture

The project follows a hexagonal (ports & adapters) architecture organized by
vertical business modules under `internal/` (`health`, `user`, `auth`,
`glossary`, `task`, and more to come: `time-registry`), plus a `shared`
module for cross-cutting code. Full conventions and rules are documented in
[CLAUDE.md](CLAUDE.md) and `.claude/rules/`.

Relevant architectural decisions are recorded as ADRs in
[documentation/adr/](documentation/adr/):

- [0001 - Hexagonal architecture organized by business module](documentation/adr/0001-hexagonal-architecture-by-module.md)
- [0002 - Custom reflection-based dependency injection container](documentation/adr/0002-custom-reflection-based-di-container.md)
- [0003 - Postgres access via pgx and a plain SQL schema](documentation/adr/0003-postgres-access-via-pgx-and-plain-sql-schema.md) (superseded by 0007)
- [0004 - JWT authentication consuming the user module's api](documentation/adr/0004-jwt-authentication-consuming-user-api.md)
- [0005 - Unified domain error structure](documentation/adr/0005-unified-domain-error-structure.md)
- [0006 - Linux-first development environment](documentation/adr/0006-linux-first-development-environment.md)
- [0007 - GORM as the project's ORM](documentation/adr/0007-gorm-as-project-orm.md)

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

This starts the app on `:8080`, a Postgres 16 instance, and
[Adminer](https://www.adminer.org/) on `:8081` to browse the database. The
schema is created by GORM's `AutoMigrate` on startup (see
[ADR 0007](documentation/adr/0007-gorm-as-project-orm.md)).

To browse the database, open `http://localhost:8081` and log in with:

- System: `PostgreSQL`
- Server: `postgres`
- Username: `tireg`
- Password: `tireg`
- Database: `tireg`

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

# use the access_token from the login response below
curl http://localhost:8080/api/v1/glossary \
  -H "Authorization: Bearer <access_token>"
```

Every `/api/v1/glossary/*` endpoint requires that `Authorization: Bearer
<token>` header and only ever reads/writes the authenticated user's own
glossary types and projects — a fresh user's first `GET
/api/v1/glossary/types` (or `GET /api/v1/glossary`) seeds 5 default types.

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

## Deploying

Production runs the API on [Render](https://render.com) (native Go
runtime, no Docker) against a [Supabase](https://supabase.com) Postgres
database, deployed via GitHub Actions rather than Render's push-triggered
auto-deploy — [.github/workflows/ci.yml](.github/workflows/ci.yml) must
pass on `main` before [.github/workflows/deploy.yml](.github/workflows/deploy.yml)
fires the Render deploy hook. No application code differs between local
Docker and production — `config.go` reads everything from env vars,
including `POSTGRES_SSLMODE` (`require` in production, `disable` locally),
and the schema is created by GORM's `AutoMigrate` on first boot, same as
locally (see [ADR 0007](documentation/adr/0007-gorm-as-project-orm.md)).

One-time setup:

1. **Supabase**: grab the direct-connection details from
   Settings → Database (host, port `5432`, user, password, database name).
   Use the direct connection, not the pooler — Render is a persistent
   service, not serverless, so it doesn't need PgBouncer.
2. **Render**: "New +" → "Blueprint", point it at this repo so it picks up
   [render.yaml](render.yaml). It creates the service without deploying
   (`autoDeploy: false` — deploys are GitHub Actions' job).
3. In the new Render service's **Environment** tab, fill in the variables
   marked `sync: false` in `render.yaml` (the Supabase values from step 1,
   plus a strong `JWT_SECRET`).
4. In the Render service's **Settings → Deploy Hook**, copy the hook URL.
5. In GitHub, repo **Settings → Secrets and variables → Actions**, add it
   as a secret named `RENDER_DEPLOY_HOOK_URL`.
6. Push to `main` (or merge a PR into it). Once CI is green, the deploy
   workflow fires the hook; watch the Render logs for the `AutoMigrate`
   table creation followed by `"starting tireg server"`.
