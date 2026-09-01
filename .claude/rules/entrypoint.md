---
paths:
  - "cmd/**"
---

# Entrypoint (`cmd/`)

Cada binario ejecutable vive en `cmd/<binario>/main.go` (p. ej.
`cmd/tireg/main.go`). El `main.go` es el único lugar que:

1. Importa el paquete `application/bootstrap` de cada módulo de negocio,
   con un alias de import igual al nombre del módulo (p. ej.
   `health "github.com/gerarc/tireg/internal/health/application/bootstrap"`).
2. Para cada módulo, llama únicamente a `modulo.WireRoutes(mux)` — nunca a
   sus funciones `Wire*` internas (`WirePersistenceDependencies`,
   `WireUseCaseDependencies`, etc.), que solo le importan al propio módulo.
   `cmd` decide en qué orden se wirean los módulos entre sí (si hay
   dependencias entre ellos), pero no el orden interno de cada uno.
3. Arranca el servidor a través del paquete de infraestructura compartido
   (`internal/shared/infrastructure/in/rest`), sin conocer detalles de
   routers o controladores concretos de ningún módulo.

```go
func wire() *http.ServeMux {
    mux := http.NewServeMux()

    health.WireRoutes(mux)

    return mux
}
```

`main.go` no contiene lógica de negocio ni construye respuestas HTTP
directamente; si algo no es "orquestar el wiring de alto nivel y arrancar el
proceso", no va en `cmd`.
