# Entorno de desarrollo

El desarrollo de este proyecto se hace originalmente en Linux.
Si estás trabajando desde Windows, no ejecutes `go build`, `go vet`,
`go test`, `docker compose`, etc. directamente en PowerShell/cmd — hazlo
dentro de WSL.

Motivo: en Windows nativo, políticas de Application Control pueden bloquear
la ejecución de los binarios que `go test` compila y corre en un directorio
temporal (`An Application Control policy has blocked this file`), lo que da
falsos negativos en paquetes que sí compilan y pasan correctamente. WSL evita
ese problema porque corre un Linux real.
