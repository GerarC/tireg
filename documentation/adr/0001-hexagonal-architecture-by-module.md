# 0001 - Hexagonal architecture organized by business module

## Status

Accepted

## Context

The project needs an architecture that keeps business logic independent from
frameworks and infrastructure details (HTTP, Postgres, queues), and that
scales to several independent business modules (`project`, `glossary`,
`time-registry`, plus shared cross-cutting code) without them becoming
tangled together.

## Decision

Adopt a hexagonal (ports & adapters) architecture, organized by **vertical
business modules** rather than by global horizontal layers. Each module
under `internal/<module>/` is self-contained and follows the same internal
structure:

```
internal/<module>/
├── application/bootstrap/        # dependency wiring for the module
├── domain/
│   ├── api/                      # inbound ports: use case interfaces
│   ├── model/                    # pure domain entities
│   ├── spi/                      # outbound ports: infrastructure interfaces
│   └── usecase/                  # business logic, implements domain/api
└── infrastructure/
    ├── in/rest/v1/                # inbound adapter: REST controllers, dto, mapper
    └── out/postgres/              # outbound adapter: repository, adapter, entity, mapper
```

`domain` never imports from `infrastructure`. `usecase` depends only on
`domain/spi` interfaces, never on concrete adapters. A `shared` module
follows the same structure and hosts cross-cutting code (e.g. the DI
container) used by more than one business module.

Full rules are documented in `CLAUDE.md`.

## Consequences

- Business logic can be tested and reasoned about without a database or an
  HTTP server.
- Adding a new module means replicating a known, consistent folder
  structure, which lowers onboarding cost.
- Swapping an adapter (e.g. Postgres for another store, REST for gRPC) does
  not require touching `domain`.
- Requires discipline to keep modules from importing each other's internals
  directly instead of going through `shared` or a module's public `api`.
