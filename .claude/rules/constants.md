---
paths:
  - "internal/**"
---

# Constantes y valores hardcodeados

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
