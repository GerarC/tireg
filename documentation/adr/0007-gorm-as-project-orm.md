# 0007 - GORM as the project's ORM

## Status

Accepted. Supersedes [ADR 0003](0003-postgres-access-via-pgx-and-plain-sql-schema.md).

## Context

ADR 0003 chose `pgx` directly with plain SQL and numbered init scripts
under `db/init/`, reasoning that the project didn't need more tooling than
`user` (its only persisted module at the time) required. Building the
per-user `glossary` CRUD module surfaced that assumption as wrong: the
project always intended to use an ORM, not hand-written SQL and manual
`Scan(...)` calls repeated per repository. Continuing with plain SQL would
mean writing that boilerplate again for every new module going forward.

## Decision

Adopt `gorm.io/gorm` with `gorm.io/driver/postgres` as the project's single
ORM and persistence layer, replacing `pgx` used directly. This is a full
replacement of ADR 0003, not an addition alongside it:

- The shared Postgres client (`internal/shared/infrastructure/out/postgres.GetClient`)
  now returns `*gorm.DB` instead of `*pgxpool.Pool`, still registered once
  in the dependency injection container and resolved by any module that
  needs it.
- Postgres entities (`infrastructure/out/postgres/entity`) carry GORM
  struct tags (`gorm:"..."`) instead of being tag-free structs mapped
  positionally in hand-written `Scan(...)` calls. They embed a shared
  `AuditEntity` (`internal/shared/infrastructure/out/postgres/entity`) for
  the `created_at`/`created_by`/`updated_at`/`updated_by` columns.
- Schema is created by GORM's `AutoMigrate`, run once per module as part of
  that module's own `WirePersistenceDependencies` (each module owns
  migrating its own tables — consistent with modules being autocontained).
  The numbered `db/init/*.sql` scripts and their `docker-compose.yml`
  mount are removed; there is no separate schema-management mechanism to
  keep in sync with the Go structs anymore.
- Repositories keep the CQRS split (command/query) already established;
  they now call GORM methods (`Create`, `Where(...).First(...)`,
  `Model(...).Count(...)`, `Updates`, `Delete`) instead of building SQL
  strings by hand. Table/constraint names that remain literal values still
  live in `infrastructure/out/postgres/util/constant`, per the project's
  "no hardcoded literals" rule.

## Consequences

- Adding a new persisted module means defining GORM-tagged entities and
  calling `AutoMigrate` for them — no new `db/init/NNN_*.sql` file, and no
  manual `Scan(...)` column-order bookkeeping.
- `AutoMigrate` only adds/adjusts columns and indexes; it does not drop or
  rename them. When the project needs real down-migrations or renames
  against databases with existing data, a migration tool
  (e.g. `gorm.io/gen`'s migration support, or a dedicated migration
  library) should be introduced and this decision revisited.
- Unique-constraint-violation handling in adapters still inspects the
  underlying `*pgconn.PgError` (the `postgres` GORM driver wraps the same
  pgx errors), so that part of ADR 0003's design is preserved, not
  discarded.
- The existing `user` module was migrated to GORM at the same time this
  ADR was adopted, so the codebase does not carry two competing
  persistence styles.
