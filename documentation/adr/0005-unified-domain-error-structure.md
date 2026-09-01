# 0005 - Unified domain error structure

## Status

Accepted

## Context

Before the `user`/`auth` modules, the project had no formal convention for
representing and reporting errors. A single structure was needed for every
error in the project, with room for multiple validation failures to be
reported at once instead of failing on the first one.

## Decision

A single struct, `internal/shared/domain/exception.DomainError`, is used for
every error in the project:

```go
type DomainError struct {
    Code      int
    Message   string
    Details   []string
    Timestamp time.Time
}
```

`Code` is the HTTP status code the error maps to (e.g. `401`, `409`, `400`).
`Message` is a generic machine-readable slug (e.g. `"AUTH_INVALID"`).
`Details` holds one human-readable string per concrete violation — a single
error can carry several details at once (e.g. registering with a username
**and** an email that are both already taken returns one `409` error with
two entries in `Details`, instead of failing on the first conflict found).

Each module defines its own error constructors in `<module>/domain/exception/`,
building on `shared`'s `DomainError` with `Code`/`Message`/detail strings
declared as constants in `<module>/domain/util/constant/`. No module uses
plain `errors.New(...)` sentinels for domain errors.

Validation (`user.Register`, `auth.Login`) accumulates every violation
before returning, rather than stopping at the first one, so a caller with
multiple invalid fields gets all of them back in one response.

The REST layer has one shared contract for error responses
(`internal/shared/infrastructure/in/rest.ErrorResponseDTO` +
`WriteError`): it extracts the `*DomainError` from any returned `error` via
`errors.As`, and writes `w.WriteHeader(domainError.Code)` directly — no
controller needs a switch statement mapping its own error codes to HTTP
statuses, because the domain error already carries the status. An error
that isn't a `*DomainError` (an unexpected/infrastructure failure) is logged
server-side and reported as a generic `500 INTERNAL_ERROR`, without leaking
internal details to the client.

## Consequences

- Domain errors encode an HTTP status directly, which is a deliberate
  coupling to a transport-layer concept. This is acceptable while REST is
  the project's only transport; if a second transport (e.g. gRPC) is added
  later, this decision should be revisited, since gRPC has its own status
  code space.
- Adding a new domain error is always the same shape: a constant pair
  (`Code`/`Message`) plus a constructor function in `domain/exception`,
  which keeps error handling consistent across modules without a shared
  registry or code-generation step.
- Because validation accumulates all violations, callers can no longer
  assume "one request, one error" — clients must be able to render a list
  of `Details`.
