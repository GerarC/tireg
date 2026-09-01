---
paths:
  - "docker-compose.yml"
  - "Dockerfile"
---

# Infraestructura local (Docker)

`docker-compose.yml` en la raíz levanta la app (`Dockerfile` multi-stage) y
Postgres para desarrollo local: `docker compose up --build`. El schema se
crea con `AutoMigrate` de GORM al arrancar, no con scripts SQL (ver
[ADR 0007](../../documentation/adr/0007-gorm-as-project-orm.md), que
reemplaza al [ADR 0003](../../documentation/adr/0003-postgres-access-via-pgx-and-plain-sql-schema.md)).
Cada módulo migra sus propias tablas como parte de su propio
`WirePersistenceDependencies` (ver `rules/wiring.md`) — una tabla nueva es
una entidad GORM nueva en ese módulo, no un archivo en `db/`.
