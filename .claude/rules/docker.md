---
paths:
  - "docker-compose.yml"
  - "Dockerfile"
  - "db/**"
---

# Infraestructura local (Docker)

`docker-compose.yml` en la raíz levanta la app (`Dockerfile` multi-stage) y
Postgres para desarrollo local: `docker compose up --build`. El schema se
crea con scripts SQL planos y numerados en `db/init/*.sql`, montados en
`/docker-entrypoint-initdb.d/` (ver
[ADR 0003](../../documentation/adr/0003-postgres-access-via-pgx-and-plain-sql-schema.md)).
Una tabla nueva es un archivo `db/init/NNN_*.sql` nuevo, coherente con el
mismo patrón de numeración incremental que las migraciones.
