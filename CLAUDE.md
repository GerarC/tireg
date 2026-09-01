# CLAUDE.md

Guía de arquitectura y convenciones para este proyecto Go. Aplica a cualquier
módulo de negocio nuevo o existente dentro de `internal/`.

## Arquitectura

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

## Estructura de cada módulo

```
internal/<modulo>/
├── application/
│   └── bootstrap/                 # wiring del módulo (ver sección Wiring)
├── domain/
│   ├── api/                       # puertos entrantes: interfaces de casos de uso
│   ├── model/                     # entidades de dominio puras (sin tags de DB/JSON)
│   ├── spi/                       # puertos salientes: interfaces hacia infraestructura
│   └── usecase/                   # implementación de la lógica de negocio (implementa domain/api)
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

## Contenedor de inyección de dependencias

Vive en `internal/shared/application/utils/container/` y es propio del
proyecto (basado en reflection y generics de Go, sin librerías externas).

- `container.GetInstance()` devuelve el singleton global del contenedor.
- `Register(constructor, optionalName ...string)` registra un constructor.
  El ciclo de vida (`SINGLETON` por defecto, `TRANSIENT`, `REQUEST`, etc.) se
  infiere del **nombre de la función constructora**: si contiene "Transient"
  se registra como `TRANSIENT`, si contiene "Request" como `REQUEST`. Nombra
  los constructores pensando en esto (p. ej. `NewUserPersistence` →
  singleton, `NewRequestScopedX` → request).
- Un constructor es una función que recibe 0 o más dependencias y devuelve
  `(T)` o `(T, error)`.
- Cuando un constructor necesita varias dependencias con nombre, agrúpalas en
  un struct de parámetros y usa el tag `inject:"nombreRegistrado"` en cada
  campo. El contenedor resuelve cada campo recursivamente.
- `container.MustResolve[T](container)` y
  `container.MustResolveNamed[T](container, "nombre")` resuelven instancias
  con generics; hacen panic si el tipo no está registrado o falla la
  construcción — están pensados para usarse durante el wiring/bootstrap, no
  en runtime de negocio dentro de un usecase.
- Registra siempre contra la interfaz (el tipo de retorno del constructor
  debe ser la interfaz `domain/api` o `domain/spi`, no la implementación
  concreta), para que el resto del sistema dependa de la abstracción.

## Wiring (`application/bootstrap`)

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

## Entrypoint (`cmd/`)

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

## Configuración y variables de entorno

Todo valor configurable por ambiente (puerto del servidor, credenciales y DSN
de base de datos, URLs de servicios externos, feature flags, etc.) se lee
como variable de entorno a través del paquete centralizado
`internal/shared/application/utils/config/`, nunca con `os.Getenv` disperso
por otros paquetes.

- `config.Load()` es el único punto de entrada: carga un archivo `.env` local
  si existe (vía `config.LoadEnvFile`, sin sobrescribir variables ya
  definidas en el entorno real del sistema — el entorno real siempre gana
  sobre el `.env`) y arma un struct `Config` con los valores resultantes,
  aplicando defaults cuando la variable no está definida.
- Los nombres de variable de entorno y sus valores por defecto son
  constantes en `internal/shared/application/utils/config/util/constant/`
  (p. ej. `PortEnvVar`, `DefaultPort`). Nunca se escribe el nombre de una
  variable de entorno como string literal fuera de ese archivo.
- `cmd/<binario>/main.go` llama a `config.Load()` una sola vez al arrancar y
  pasa los valores necesarios (p. ej. el puerto) a las funciones de
  infraestructura correspondientes (`rest.HttpRestAPIInitializer(mux, appConfig.Port)`).
  Ningún paquete de infraestructura lee variables de entorno por su cuenta.
- Cuando se agregue una nueva fuente de configuración (p. ej. credenciales de
  Postgres), se extiende el mismo struct `Config` y las mismas constantes de
  `config/util/constant/`, siguiendo el patrón ya establecido — no se crea un
  mecanismo de configuración paralelo por módulo.
- `.env` está en `.gitignore` y nunca se commitea; `.env.example` en la raíz
  documenta las variables esperadas con valores de ejemplo (sin secretos
  reales) y se mantiene actualizado cada vez que se agrega una variable
  nueva.

## Constantes y valores hardcodeados

Ningún valor literal (strings, números mágicos, keys de configuración,
nombres de rutas HTTP, mensajes de error, nombres de tablas/columnas, códigos
de tipo, etc.) va escrito directamente en la lógica. Todo valor hardcodeado
se declara como constante en un archivo dentro de una carpeta `util/constant`
propia de la capa donde se usa:

```
domain/util/constant/                        # constantes de dominio (nombres de reglas, códigos de error de negocio, etc.)
application/util/constant/                   # constantes de wiring/bootstrap (nombres registrados en el container, etc.)
infrastructure/in/rest/v1/util/constant/     # constantes de la capa REST (rutas, headers, mensajes de respuesta)
infrastructure/out/postgres/util/constant/   # constantes de la capa de persistencia (nombres de tabla/columna, queries fijas)
```

Generaliza esta misma regla a cualquier otra carpeta de `infrastructure/out`
que se agregue (p. ej. `infrastructure/out/redis/util/constant`,
`infrastructure/out/sqs/util/constant`): cada adapter de salida y cada
entrada define sus propias constantes, no las comparte con otro adapter salvo
que el valor viva en `shared`.

Antes de escribir un literal en código, pregúntate a qué capa pertenece ese
valor y colócalo en el `util/constant` correspondiente a esa capa dentro del
módulo. Si el valor es transversal a todo el proyecto (no a un módulo o capa
en particular), va en `internal/shared/.../util/constant`.

## Comentarios y documentación

- No se escriben comentarios en el código de implementación (nada de `//`
  explicando qué hace una línea o un bloque). El código debe ser legible por
  sí mismo mediante nombres claros.
- Toda interfaz (`domain/api`, `domain/spi`, y cualquier otra interfaz
  pública del módulo) debe llevar un comentario de documentación en formato
  Go doc (comentario `//` inmediatamente antes de la declaración, empezando
  con el nombre del identificador) describiendo su propósito y contrato.
  Esta es la única excepción a "no comentarios": documentar interfaces es
  obligatorio, comentar implementaciones no está permitido.

## Documentación de endpoints (Swagger)

Todo endpoint expuesto en `infrastructure/in/rest/v1/controller` debe llevar
anotaciones Swagger (formato `swaggo/swag`: comentarios `// @...` inmediatamente
antes de la función del handler) describiendo al menos: resumen, descripción,
método y ruta, tags, parámetros de entrada, y las respuestas posibles con su
código HTTP y el `dto` asociado. Esto es, junto con la documentación de
interfaces, la otra excepción a "no comentarios en el código": todo endpoint
sin su bloque Swagger se considera incompleto.

## Testing

Los tests viven fuera de `internal/`, en una carpeta `test/` en la raíz del
proyecto que **replica la misma estructura de carpetas** del código que
prueban:

```
test/<modulo>/domain/usecase/x_usecase_test.go
test/<modulo>/infrastructure/in/rest/v1/mapper/x_mapper_test.go
test/<modulo>/infrastructure/in/rest/v1/controller/x_controller_test.go
test/shared/application/utils/container/container_test.go
```

Reglas:

- Cada archivo de test usa un paquete "black-box" con sufijo `_test`
  (`usecase_test`, `mapper_test`, `controller_test`, etc.) e importa el
  paquete bajo prueba por su ruta de módulo completa
  (`github.com/gerarc/tireg/internal/<modulo>/...`). No se usan paquetes
  internos ni acceso a no exportados: se prueba el contrato público.
- Solo `testing` y `net/http/httptest` de la librería estándar; no se agrega
  un framework de testing/mocking externo sin necesidad real.
- Cuando una dependencia es una interfaz de `domain/api` o `domain/spi`, se
  implementa un stub mínimo en el propio archivo de test en vez de traer un
  generador de mocks.
- Un test por comportamiento relevante del componente (camino feliz, casos
  límite, errores esperados); nombres de test descriptivos en inglés
  (`TestX_DoesY_WhenZ`).
- Solo se crean las carpetas de `test/` que correspondan a código que
  realmente existe — igual que con la plantilla de `internal/`, no se
  rellena la estructura por adelantado.

## Documentación arquitectónica (ADR)

Las decisiones arquitectónicas relevantes (elegir un patrón de arquitectura,
construir el contenedor DI propio, adoptar o descartar una librería/base de
datos, cambios de convención que afectan a todo el proyecto, etc.) se
documentan como Architecture Decision Records en `documentation/adr/`,
numerados secuencialmente: `NNNN-titulo-en-kebab-case.md`.

Cada ADR sigue el formato estándar: `Status`, `Context`, `Decision`,
`Consequences`. Se escriben en inglés, igual que el resto del código y su
documentación. Un ADR nuevo puede referenciar a uno anterior con un link
relativo (p. ej. `[ADR 0001](0001-titulo.md)`) cuando construye sobre una
decisión previa. Los ADR no se editan retroactivamente para cambiar la
decisión tomada — si una decisión se revierte o reemplaza, se crea un ADR
nuevo que la reemplaza y se referencia mutuamente.

## Idioma

Todo el código debe escribirse en inglés: nombres de paquetes, tipos,
funciones, variables, constantes, archivos, mensajes de error y la
documentación de las interfaces. Esto aplica sin excepción, incluso cuando el
dominio de negocio o los datos de origen (por ejemplo el registro de horas)
estén en español — la traducción al inglés ocurre al modelar el dominio.

## Convenciones generales

- Nombres de paquete en inglés y en singular salvo que el dominio use
  naturalmente el plural; nombres de módulo/carpeta pueden ir en kebab-case
  si tienen más de una palabra (p. ej. `time-registry`).
- No agregues dependencias externas (frameworks web, ORMs) sin necesidad real
  del caso de uso que se está implementando — evalúa primero si el estándar
  library alcanza, y si no, elige una sola herramienta por responsabilidad
  (un router, un driver/ORM de Postgres) en vez de mezclar varias.
- No implementes capas o carpetas que un módulo no necesita todavía (p. ej.
  no crear `infrastructure/out/postgres` en un módulo que no persiste nada
  por sí mismo). La plantilla de carpetas es una guía, no una obligación de
  llenarla toda.
- Mantén `domain/model` libre de tags de serialización/persistencia; esos
  tags viven en `dto` (REST) y `entity` (Postgres) respectivamente.
