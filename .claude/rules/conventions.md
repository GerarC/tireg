# Idioma

Todo el código debe escribirse en inglés: nombres de paquetes, tipos,
funciones, variables, constantes, archivos, mensajes de error y la
documentación de las interfaces. Esto aplica sin excepción, incluso cuando el
dominio de negocio o los datos de origen (por ejemplo el registro de horas)
estén en español — la traducción al inglés ocurre al modelar el dominio.

# Convenciones generales

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
