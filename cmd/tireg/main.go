package main

import (
	"log"
	"net/http"

	_ "github.com/gerarc/tireg/docs"

	auth "github.com/gerarc/tireg/internal/auth/application/bootstrap"
	glossary "github.com/gerarc/tireg/internal/glossary/application/bootstrap"
	health "github.com/gerarc/tireg/internal/health/application/bootstrap"
	shared "github.com/gerarc/tireg/internal/shared/application/bootstrap"
	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/application/utils/container"
	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	user "github.com/gerarc/tireg/internal/user/application/bootstrap"
)

func wire() *http.ServeMux {
	mux := http.NewServeMux()

	health.WireRoutes(mux)
	user.WireRoutes(mux)
	auth.WireRoutes(mux)
	glossary.WireRoutes(mux)

	rest.RegisterSwaggerRoutes(mux)

	return mux
}

// @title tireg API
// @version 1.0
// @description Backend for time registration, projects, and glossary management.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT access token returned by /api/v1/auth/login.
func main() {
	shared.WireInfrastructureDependencies()

	appContainer := container.GetInstance()
	appConfig := container.MustResolve[*config.Config](appContainer)

	mux := wire()

	log.Println("Starting tireg server...")

	if err := rest.HttpRestAPIInitializer(mux, appConfig.Port); err != nil {
		log.Fatalf("critical error starting server: %v", err)
	}
}
