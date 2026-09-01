---
paths:
  - "documentation/adr/**"
---

# Documentación arquitectónica (ADR)

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
