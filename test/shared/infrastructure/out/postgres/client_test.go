package postgres_test

import (
	"net/url"
	"testing"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres"
)

func TestBuildDSN_ProducesParsableURL_WhenPasswordHasSpecialCharacters(t *testing.T) {
	appConfig := &config.Config{
		PostgresUser:     "postgres",
		PostgresPassword: "p@ss:w/ord?with#special&chars",
		PostgresHost:     "db.example.supabase.co",
		PostgresPort:     "5432",
		PostgresDBName:   "postgres",
		PostgresSSLMode:  "require",
	}

	dsn := postgres.BuildDSN(appConfig)

	parsed, err := url.Parse(dsn)
	if err != nil {
		t.Fatalf("expected a parsable DSN, got error: %v (dsn=%q)", err, dsn)
	}

	if parsed.Hostname() != "db.example.supabase.co" || parsed.Port() != "5432" {
		t.Fatalf("expected host db.example.supabase.co:5432, got %s:%s", parsed.Hostname(), parsed.Port())
	}

	password, _ := parsed.User.Password()
	if password != appConfig.PostgresPassword {
		t.Fatalf("expected password to round-trip unchanged, got %q", password)
	}

	if parsed.Query().Get("sslmode") != "require" {
		t.Fatalf("expected sslmode=require, got %q", parsed.Query().Get("sslmode"))
	}
}

func TestBuildDSN_ProducesExpectedFormat_WithPlainCredentials(t *testing.T) {
	appConfig := &config.Config{
		PostgresUser:     "tireg",
		PostgresPassword: "tireg",
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		PostgresDBName:   "tireg",
		PostgresSSLMode:  "disable",
	}

	dsn := postgres.BuildDSN(appConfig)

	expected := "postgres://tireg:tireg@localhost:5432/tireg?sslmode=disable"
	if dsn != expected {
		t.Fatalf("expected %q, got %q", expected, dsn)
	}
}
