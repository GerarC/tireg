package config_test

import (
	"os"
	"testing"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
)

func TestLoad_UsesDefaultPortWhenEnvVarNotSet(t *testing.T) {
	os.Unsetenv("PORT")

	appConfig := config.Load()

	if appConfig.Port != "8080" {
		t.Fatalf("expected default port %q, got %q", "8080", appConfig.Port)
	}
}

func TestLoad_UsesPortEnvVarWhenSet(t *testing.T) {
	os.Setenv("PORT", "9090")
	t.Cleanup(func() {
		os.Unsetenv("PORT")
	})

	appConfig := config.Load()

	if appConfig.Port != "9090" {
		t.Fatalf("expected port %q, got %q", "9090", appConfig.Port)
	}
}
