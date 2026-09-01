# 0003 - Postgres access via pgx and a plain SQL schema

## Status

Accepted

## Context

The `user` module is the first module in the project that needs real
persistence. A Postgres driver and a way to create the database schema had
to be chosen, without adopting more tooling than the project currently
needs.

## Decision

Use `github.com/jackc/pgx/v5` directly as the project's only Postgres
driver, through a single shared connection pool
(`internal/shared/infrastructure/out/postgres.GetClient`, `*pgxpool.Pool`)
registered once in the dependency injection container via
`internal/shared/application/bootstrap.WireInfrastructureDependencies`. Any
module that needs the pool declares it as a constructor parameter; the
container resolves the same singleton instance.

No ORM and no query builder are introduced — repositories write plain SQL
queries as constants in their own `infrastructure/out/postgres/util/constant`
package, following the project's "no hardcoded literals" rule.

The database schema is created with plain, numbered SQL files under
`db/init/` (e.g. `001_create_users_table.sql`), mounted by `docker-compose`
into Postgres' `/docker-entrypoint-initdb.d/`, which Postgres executes
automatically on first startup. No migration library
(e.g. `golang-migrate`) is introduced yet.

## Consequences

- Adding a new table means adding a new `db/init/NNN_*.sql` file, but this
  only provisions a fresh database volume — it does not apply incremental
  migrations against an already-running database with existing data.
- Unique constraint names created implicitly by Postgres (e.g.
  `users_username_key`) are relied upon by adapters to translate unique
  violations into domain errors, which is an implicit coupling between the
  SQL schema and the Go code that must be kept in sync manually.
- When the project needs real migrations against existing databases (not
  just fresh local/dev environments), a migration library should be
  introduced and this decision revisited.
