# 0004 - JWT authentication consuming the user module's api

## Status

Accepted

## Context

A login flow was needed on top of the existing user registration, without
duplicating access to user data in a second module.

## Decision

`internal/auth/` has no persistence of its own. Its use case
(`auth/domain/usecase.AuthUseCaseImplemented`) depends directly on
`user/domain/api.UserUseCase` (which gains a `FindByIdentifier` method),
using the exception `CLAUDE.md` already allows for a module consuming
another module's public `api`.

Access tokens are signed JWTs (`github.com/golang-jwt/jwt/v5`, HS256) issued
through the port `auth/domain/spi.TokenIssuer`, implemented in
`auth/infrastructure/out/jwt/adapter` using a secret and expiration read
from the shared `config` package. Login accepts a single `identifier` field
(username or email), since the repository already resolves both with one
query. Password hashing/verification is delegated to the shared
`PasswordHasher` port (`internal/shared/domain/spi`), the first port shared
by two business modules, which is exactly the case `CLAUDE.md` already
describes for putting something in `shared`.

To avoid user enumeration, a missing user and a wrong password both produce
the same generic `AUTH_INVALID` error. Missing/empty `identifier` or
`password` are treated as a separate `AUTH_VALIDATION_FAILED` (400) error,
accumulating all missing-field details at once, following the same
validation pattern used by `user.Register`.

No middleware, no protected routes, and no refresh tokens are introduced in
this pass — the scope is strictly registration + login.

## Consequences

- `auth` has a compile-time dependency on `user`'s `domain/api` package;
  `cmd` must wire `user` before `auth` so `UserUseCase` is registered in the
  container by the time `auth`'s use case is resolved.
- Tokens are stateless and cannot be individually revoked without an
  additional deny-list mechanism, which does not exist yet.
- Adding protected endpoints later requires a new piece (token validation
  middleware), deliberately left out of this decision.
