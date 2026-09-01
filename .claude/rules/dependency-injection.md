---
paths:
  - "internal/**"
---

# Contenedor de inyección de dependencias

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
- No toda interfaz registrada es un puerto de dominio: utilidades
  transversales que ningún caso de uso necesita por inversión de
  dependencias (p. ej. `logger.Logger` en
  `internal/shared/application/utils/logger/`) se definen junto a
  `config`/`container` bajo `application/utils/`, no en `domain/spi`. Se
  registran igual que cualquier otro constructor (`container.Register(logger.NewSlogAdapter)`)
  y, al no tener "Transient"/"Request" en el nombre, el container ya las
  trata como `SINGLETON` por defecto — una sola instancia cacheada tras el
  primer `MustResolve`, sin necesitar un singleton global aparte.
