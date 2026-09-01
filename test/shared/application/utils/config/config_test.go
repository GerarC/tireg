package config_test

import (
	"os"
	"testing"
	"time"

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

func TestLoad_UsesPostgresDefaults_WhenEnvVarsNotSet(t *testing.T) {
	for _, envVar := range []string{"POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB", "POSTGRES_SSLMODE"} {
		os.Unsetenv(envVar)
	}

	appConfig := config.Load()

	if appConfig.PostgresHost != "localhost" {
		t.Fatalf("expected default postgres host %q, got %q", "localhost", appConfig.PostgresHost)
	}

	if appConfig.PostgresPort != "5432" {
		t.Fatalf("expected default postgres port %q, got %q", "5432", appConfig.PostgresPort)
	}

	if appConfig.PostgresSSLMode != "disable" {
		t.Fatalf("expected default postgres sslmode %q, got %q", "disable", appConfig.PostgresSSLMode)
	}
}

func TestLoad_UsesPostgresEnvVarsWhenSet(t *testing.T) {
	os.Setenv("POSTGRES_HOST", "db.internal")
	os.Setenv("POSTGRES_PASSWORD", "super-secret")
	t.Cleanup(func() {
		os.Unsetenv("POSTGRES_HOST")
		os.Unsetenv("POSTGRES_PASSWORD")
	})

	appConfig := config.Load()

	if appConfig.PostgresHost != "db.internal" {
		t.Fatalf("expected postgres host %q, got %q", "db.internal", appConfig.PostgresHost)
	}

	if appConfig.PostgresPassword != "super-secret" {
		t.Fatalf("expected postgres password %q, got %q", "super-secret", appConfig.PostgresPassword)
	}
}

func TestLoad_UsesDefaultJWTExpiration_WhenEnvVarNotSet(t *testing.T) {
	os.Unsetenv("JWT_EXPIRATION_MINUTES")

	appConfig := config.Load()

	if appConfig.JWTExpiration != 60*time.Minute {
		t.Fatalf("expected default jwt expiration of 60 minutes, got %v", appConfig.JWTExpiration)
	}
}

func TestLoad_UsesJWTExpirationEnvVarWhenSet(t *testing.T) {
	os.Setenv("JWT_EXPIRATION_MINUTES", "15")
	t.Cleanup(func() {
		os.Unsetenv("JWT_EXPIRATION_MINUTES")
	})

	appConfig := config.Load()

	if appConfig.JWTExpiration != 15*time.Minute {
		t.Fatalf("expected jwt expiration of 15 minutes, got %v", appConfig.JWTExpiration)
	}
}

func TestLoad_UsesDefaultLogLevel_WhenEnvVarNotSet(t *testing.T) {
	os.Unsetenv("LOG_LEVEL")

	appConfig := config.Load()

	if appConfig.LogLevel != "info" {
		t.Fatalf("expected default log level %q, got %q", "info", appConfig.LogLevel)
	}
}

func TestLoad_UsesLogLevelEnvVarWhenSet(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	t.Cleanup(func() {
		os.Unsetenv("LOG_LEVEL")
	})

	appConfig := config.Load()

	if appConfig.LogLevel != "debug" {
		t.Fatalf("expected log level %q, got %q", "debug", appConfig.LogLevel)
	}
}
