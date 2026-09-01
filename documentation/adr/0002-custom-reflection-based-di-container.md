# 0002 - Custom reflection-based dependency injection container

## Status

Accepted

## Context

The hexagonal architecture from [ADR 0001](0001-hexagonal-architecture-by-module.md)
requires wiring concrete adapters and use cases behind interfaces
(`domain/api`, `domain/spi`) at application startup, across several
independent modules. A mechanism was needed to register constructors and
resolve fully-built instances (with their own dependencies resolved
recursively) without each module having to manually wire every dependency of
every other module by hand.

## Decision

Build a small dependency injection container from scratch
(`internal/shared/application/utils/container/`), based on Go's `reflect`
package and generics, instead of adopting an external DI library. Key
design points:

- `Register(constructor, optionalName ...string)` registers a constructor
  function. Its lifecycle (`SINGLETON` by default, `TRANSIENT`, `REQUEST`,
  `SESSION`, `APPLICATION`, `WEBSOCKET`) is inferred from the constructor's
  function name (e.g. containing `Transient` or `Request`).
- `MustResolve[T]` / `MustResolveNamed[T]` resolve an instance of type `T`
  using generics, avoiding manual type assertions at call sites.
- Dependencies declared as a struct parameter with `inject:"name"` tags are
  resolved recursively, field by field.
- A single global instance is exposed via `container.GetInstance()`
  (`sync.Once` singleton), used by each module's `application/bootstrap`
  package during wiring, and orchestrated from `cmd`.

## Consequences

- No external dependency is introduced for wiring; the container stays
  small and fully understood by the team.
- Constructors must be named with the lifecycle convention in mind
  (`Transient`/`Request` substrings), which is an implicit contract that
  must be documented and respected.
- Resolution failures (missing registration, constructor error) are only
  caught at runtime via panics, not at compile time — wiring must be
  exercised (e.g. by starting the application or a wiring test) to catch
  mistakes early.
- Adding new lifecycle scopes or resolution behavior requires modifying the
  container itself rather than configuring a third-party library.
