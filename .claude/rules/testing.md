---
paths:
  - "test/**"
---

# Testing

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
