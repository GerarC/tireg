---
paths:
  - "internal/**"
---

# Comentarios y documentación

- No se escriben comentarios en el código de implementación (nada de `//`
  explicando qué hace una línea o un bloque). El código debe ser legible por
  sí mismo mediante nombres claros.
- Toda interfaz (`domain/api`, `domain/spi`, y cualquier otra interfaz
  pública del módulo) debe llevar un comentario de documentación en formato
  Go doc (comentario `//` inmediatamente antes de la declaración, empezando
  con el nombre del identificador) describiendo su propósito y contrato.
  Esta es la única excepción a "no comentarios": documentar interfaces es
  obligatorio, comentar implementaciones no está permitido.
