package main

import (
	"log"
	"net/http"

	health "github.com/gerarc/tireg/internal/health/application/bootstrap"
	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
)

func wire() *http.ServeMux {
	mux := http.NewServeMux()

	health.WireRoutes(mux)

	return mux
}

func main() {
	appConfig := config.Load()
	mux := wire()

	log.Println("Starting tireg server...")

	if err := rest.HttpRestAPIInitializer(mux, appConfig.Port); err != nil {
		log.Fatalf("critical error starting server: %v", err)
	}
}
