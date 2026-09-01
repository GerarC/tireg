---
paths:
  - "internal/**/application/bootstrap/**"
  - "cmd/**"
---

# Wiring (`application/bootstrap`)

Cada módulo expone funciones `Wire*` en su paquete `application/bootstrap`
(o subpaquetes de este, p. ej. `bootstrap/persistence`, `bootstrap/handlers`,
si el wiring de un módulo crece) que registran sus componentes en el
contenedor global, una función por responsabilidad. Patrón:

```go
package bootstrap

import (
    "net/http"

    "github.com/gerarc/tireg/internal/shared/application/utils/container"
    "github.com/gerarc/tireg/internal/<modulo>/domain/usecase"
    "github.com/gerarc/tireg/internal/<modulo>/infrastructure/in/rest/v1/controller"
    "github.com/gerarc/tireg/internal/<modulo>/infrastructure/out/postgres/adapter"
    "github.com/gerarc/tireg/internal/<modulo>/infrastructure/out/postgres/repository"
)

func WirePersistenceDependencies() {
    c := container.GetInstance()

    c.Register(repository.NewXRepository)
    c.Register(adapter.NewXAdapter)
}

func WireUseCaseDependencies() {
    c := container.GetInstance()

    c.Register(usecase.NewXUseCase)
}

func WireHandlerDependencies() {
    c := container.GetInstance()

    c.Register(controller.NewXController)
}

func WireRoutes(mux *http.ServeMux) {
    WirePersistenceDependencies()
    WireUseCaseDependencies()
    WireHandlerDependencies()

    c := container.GetInstance()

    xController := container.MustResolve[*controller.XController](c)
    xController.RegisterRoutes(mux)
}
```

`WireRoutes(mux)` es el **único punto de entrada del módulo hacia `cmd`**:
internamente orquesta, en el orden correcto, todas las funciones `Wire*` de
ese módulo (persistencia → casos de uso → handlers) y al final registra sus
rutas en el mux. Ningún otro `Wire*` del módulo se llama directamente desde
`cmd` — solo desde dentro de `WireRoutes` (o entre sí, dentro del propio
paquete `bootstrap`). Esto mantiene el orden de arranque interno de un
módulo como un detalle propio del módulo, no algo que `cmd` deba conocer o
repetir.

Orden general a respetar dentro de `WireRoutes`: infraestructura de bajo
nivel (DB, clientes externos) → persistencia/adapters → casos de uso →
handlers/controllers → registro de rutas. Un módulo nunca debería registrar
su `usecase` antes que las dependencias `spi` que ese `usecase` necesita.

Cuando un módulo consume el `domain/api` de otro (ver `rules/architecture.md`),
`cmd` debe invocar `WireRoutes` del módulo consumido **antes** que el del
módulo consumidor — p. ej. `user.WireRoutes(mux)` antes que
`auth.WireRoutes(mux)`, porque el `usecase` de `auth` resuelve
`user/domain/api.UserUseCase` desde el container, y esa interfaz solo queda
registrada una vez que `user.WireRoutes` corrió.
