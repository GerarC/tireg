---
paths:
  - "internal/**"
---

# Arquitectura

Arquitectura hexagonal (ports & adapters) organizada por **módulos de negocio
verticales**, no por capas horizontales globales. Cada módulo vive en
`internal/<modulo>/` y es autocontenido: tiene su propio `domain`,
`application` e `infrastructure`. El módulo `internal/shared/` sigue la misma
estructura y aloja código transversal (utilidades, el contenedor DI, tipos
comunes) que varios módulos de negocio pueden necesitar.

No metas lógica de negocio de un módulo dentro de otro. Si dos módulos
necesitan compartir algo, ese algo va a `shared`, no se importa directamente
de módulo a módulo salvo que sea estrictamente necesario (p. ej. un módulo
consumiendo el `api` público de otro).

Caso real ya implementado: `internal/auth/domain/usecase` importa
`internal/user/domain/api.UserUseCase` directamente (para resolver el
usuario por `identifier` al hacer login) en vez de tener su propio acceso a
Postgres — evita duplicar acceso a los mismos datos desde dos módulos. Este
es el patrón a seguir cuando un módulo necesita datos u operaciones que ya
son responsabilidad de otro: consumir su `domain/api`, nunca su
`infrastructure` ni su `domain/spi`.

## Estructura de cada módulo

```
internal/<modulo>/
├── application/
│   └── bootstrap/                 # wiring del módulo (ver rules/wiring.md)
├── domain/
│   ├── api/                       # puertos entrantes: interfaces de casos de uso
│   ├── model/                     # entidades de dominio puras (sin tags de DB/JSON)
│   ├── spi/                       # puertos salientes: interfaces hacia infraestructura
│   ├── usecase/                   # implementación de la lógica de negocio (implementa domain/api)
│   └── exception/                 # constructores de errores de dominio (ver rules/error-handling.md)
└── infrastructure/
    ├── in/
    │   └── rest/v1/
    │       ├── controller/        # handlers HTTP, delegan al usecase vía domain/api
    │       ├── dto/                # request/response de la API
    │       ├── mapper/             # DTO <-> modelo de dominio
    │       └── util/constant/      # constantes propias de esta capa (rutas, headers, etc.)
    └── out/
        └── postgres/
            ├── adapter/            # implementa domain/spi contra Postgres
            ├── entity/              # entidades de persistencia (structs con tags de DB)
            ├── mapper/              # entidad <-> modelo de dominio
            └── repository/         # acceso a datos (queries, sin lógica de negocio)
```

### Reglas de capas

- **`domain`** no importa nada de `infrastructure`. Es el núcleo: no conoce
  Postgres, HTTP, ni ningún detalle de implementación.
- **`domain/api`** define qué puede hacer el módulo hacia afuera (los casos de
  uso como interfaces). **`domain/usecase`** implementa esas interfaces.
- **`domain/spi`** define qué necesita el módulo de infraestructura (p. ej.
  un repositorio). El `usecase` depende de la interfaz `spi`, nunca de un
  adapter concreto.
- **`infrastructure/out/*`** implementa las interfaces `spi`. Un `adapter`
  traduce entre el mundo externo (Postgres, una API externa, una cola) y el
  modelo de dominio, apoyándose en su propio `mapper` y `entity`/`dto`.
- **`infrastructure/in/*`** expone el módulo hacia afuera (REST, hoy; podría
  haber `grpc`, `graphql`, `cli`, etc. como hermanos de `rest` si se necesita).
  Un `controller` solo traduce HTTP <-> dominio vía `dto`/`mapper` y delega
  toda la lógica al `usecase` a través de `domain/api`.
- Nunca uses el `entity` de persistencia ni el `dto` de REST como si fueran el
  modelo de dominio. Siempre se mapea explícitamente.
- Versiona la API entrante (`rest/v1`, `rest/v2`, ...) cuando haya cambios
  incompatibles; no rompas `v1` para acomodar un cambio nuevo.
