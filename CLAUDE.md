# CLAUDE.md

Este proyecto es un backend Go con arquitectura hexagonal organizada por
módulos de negocio verticales (`internal/<modulo>/`).

Las reglas de arquitectura y convenciones detalladas viven en
`.claude/rules/` (se cargan automáticamente por tema; algunas están
scoped por `paths` a las carpetas donde aplican):

- [development-environment.md](.claude/rules/development-environment.md) — desarrollo en Linux; en Windows usar WSL
- [architecture.md](.claude/rules/architecture.md) — arquitectura hexagonal, estructura de módulo, reglas de capas
- [dependency-injection.md](.claude/rules/dependency-injection.md) — contenedor DI propio (`container.GetInstance`, `Register`, ciclos de vida)
- [wiring.md](.claude/rules/wiring.md) — funciones `Wire*` en `application/bootstrap`
- [entrypoint.md](.claude/rules/entrypoint.md) — convenciones de `cmd/<binario>/main.go`
- [configuration.md](.claude/rules/configuration.md) — variables de entorno vía `config.Load()`
- [constants.md](.claude/rules/constants.md) — nada de literales hardcodeados, todo en `util/constant`
- [error-handling.md](.claude/rules/error-handling.md) — `DomainError` unificado (ver [ADR 0005](documentation/adr/0005-unified-domain-error-structure.md))
- [comments-and-docs.md](.claude/rules/comments-and-docs.md) — sin comentarios salvo Go doc en interfaces
- [swagger.md](.claude/rules/swagger.md) — anotaciones `swaggo/swag` y regeneración del spec
- [testing.md](.claude/rules/testing.md) — estructura de `test/`, black-box testing
- [adr.md](.claude/rules/adr.md) — Architecture Decision Records en `documentation/adr/`
- [docker.md](.claude/rules/docker.md) — `docker-compose`, schema en `db/init/*.sql`
- [conventions.md](.claude/rules/conventions.md) — idioma (inglés en código) y convenciones generales

Si vas a agregar una regla nueva o cambiar una existente, edita el archivo
de `.claude/rules/` correspondiente al tema (o crea uno nuevo) en vez de
escribirla aquí.
