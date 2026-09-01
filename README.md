# tireg

A backend to use with models, used for maintain a local database for time registration.

## Architecture

The project follows a hexagonal (ports & adapters) architecture organized by
vertical business modules under `internal/` (`health`, and more to come:
`project`, `glossary`, `time-registry`), plus a `shared` module for
cross-cutting code. Full conventions and rules are documented in
[CLAUDE.md](CLAUDE.md).

Relevant architectural decisions are recorded as ADRs in
[documentation/adr/](documentation/adr/):

- [0001 - Hexagonal architecture organized by business module](documentation/adr/0001-hexagonal-architecture-by-module.md)
- [0002 - Custom reflection-based dependency injection container](documentation/adr/0002-custom-reflection-based-di-container.md)

## Requirements

- Go 1.25.4+

> **Windows:** run everything (build, run, test) inside **WSL**, not
> directly in PowerShell/cmd. Windows Smart App Control blocks locally
> built, unsigned `.exe` files (including `go run` and `go test` binaries),
> which makes native Windows execution unreliable. WSL binaries are plain
> Linux ELF executables and are unaffected.

## Running the server

```bash
go build -o tireg ./cmd/tireg && ./tireg
```

The server listens on `:8080` by default (override with the `PORT` env var).
A local `.env` file (see `.env.example`) is also loaded on startup, without
overriding real environment variables.

Once running, check the health endpoint:

```bash
curl http://localhost:8080/api/v1/health
```

## Testing

Tests live under `test/`, mirroring the structure of `internal/`:

```bash
go test ./...
```
