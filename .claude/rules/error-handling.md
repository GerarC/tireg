---
paths:
  - "internal/**"
---

# Manejo de errores

Todo error del proyecto usa una única estructura,
`internal/shared/domain/exception.DomainError{Code int; Message string;
Details []string; Timestamp time.Time}` (ver
[ADR 0005](../../documentation/adr/0005-unified-domain-error-structure.md)).

- **`Code`** es directamente el status HTTP (`400`, `401`, `404`, `409`,
  ...). Cada módulo declara sus propios códigos como constantes enteras en
  `domain/util/constant` (nunca importando `net/http` desde `domain`, solo
  reutilizando el mismo número).
- **`Message`** es un slug genérico en mayúsculas (`"USER_ALREADY_TAKEN"`,
  `"AUTH_INVALID"`), también constante en `domain/util/constant`.
- **`Details`** es la lista de strings humanos, uno por violación concreta.
  Nunca se corta en el primer error: si una operación tiene varias
  violaciones simultáneas (dos campos inválidos, username Y email ya
  tomados), se acumulan **todas** en una sola respuesta.
- Cada módulo expone solo *constructores* de `*exception.DomainError` en su
  propio `domain/exception/` (p. ej. `user/domain/exception.NewAlreadyTakenError(details...)`).
  Ningún módulo usa `errors.New(...)` suelto para errores de negocio.
- La capa REST tiene un contrato único en `internal/shared/infrastructure/in/rest/`:
  `ErrorResponseDTO` (mismo shape que `DomainError`, serializado a JSON) y
  `WriteError(w, err)`, que hace `errors.As` para extraer el `*DomainError`
  (o construye uno genérico `INTERNAL_ERROR`/500 si el error no lo es, sin
  filtrar detalles internos) y escribe `w.WriteHeader(domainError.Code)`
  directamente. Ningún controller necesita mapear su propio código de
  negocio a un status HTTP — el error ya lo trae. Un controller solo llama
  `sharedRest.WriteError(w, err)`.
- Cuerpos de request inválidos (JSON malformado) usan
  `sharedRest.InvalidRequestBodyError(err)`, que arma el mismo tipo de error
  (código 400) — es un error de infraestructura REST, no de dominio, pero
  usa la misma estructura envolvente para que la API sea consistente.
