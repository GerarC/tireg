---
paths:
  - "internal/**/infrastructure/in/rest/**/controller/**"
---

# Documentación de endpoints (Swagger)

Todo endpoint expuesto en `infrastructure/in/rest/v1/controller` debe llevar
anotaciones Swagger (formato `swaggo/swag`: comentarios `// @...` inmediatamente
antes de la función del handler) describiendo al menos: resumen, descripción,
método y ruta, tags, parámetros de entrada, y las respuestas posibles con su
código HTTP y el `dto` asociado. Esto es, junto con la documentación de
interfaces, la otra excepción a "no comentarios en el código": todo endpoint
sin su bloque Swagger se considera incompleto.

El spec se genera y se sirve de verdad (no son solo comentarios muertos):

- `swag` está declarado como tool dependency de Go (`tool
  github.com/swaggo/swag/cmd/swag` en `go.mod`, instalado con
  `go get -tool ...`), no como dependencia de build.
- Tras tocar cualquier anotación (o agregar un endpoint), hay que regenerar
  el spec y commitear el resultado:
  `go tool swag init -g cmd/tireg/main.go -o docs --parseDependency`.
  `--parseDependency` es obligatorio porque los controllers referencian
  `sharedRest.ErrorResponseDTO`, un tipo fuera del propio paquete del
  controller — sin ese flag, `swag` no lo resuelve.
- **Trampa conocida**: `swag` resuelve un tipo `pkgAlias.Tipo` en una
  anotación buscando el import `pkgAlias` **realmente usado** en ese mismo
  archivo `.go`. Si el handler solo usa el DTO de forma implícita (p. ej.
  `healthResponseDTO := mapper.ToHealthResponseDTO(...)` sin nombrar nunca
  el tipo), Go no obliga a importar el paquete del DTO y `swag` falla con
  "cannot find type definition". Si esto pasa, tipa la variable
  explícitamente (`var responseDTO dto.HealthResponseDTO = mapper.ToX(...)`)
  para forzar el import real — no falsifiques el import con un
  `var _ = ...` de relleno.
- `internal/shared/infrastructure/in/rest.RegisterSwaggerRoutes(mux)` monta
  la UI en `/swagger/index.html`; `cmd` la invoca una sola vez, junto a las
  `WireRoutes` de cada módulo.
