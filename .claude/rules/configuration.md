---
paths:
  - "internal/shared/application/utils/config/**"
  - "cmd/**"
  - ".env.example"
---

# Configuración y variables de entorno

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
