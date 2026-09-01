package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
)

func TestLoadEnvFile_SetsUnsetEnvironmentVariables(t *testing.T) {
	envFilePath := filepath.Join(t.TempDir(), ".env")
	envFileContent := "FOO_TEST_VAR=bar\n# a comment\n\nBAZ_TEST_VAR=qux\n"

	if err := os.WriteFile(envFilePath, []byte(envFileContent), 0o600); err != nil {
		t.Fatalf("unexpected error writing env file: %v", err)
	}

	os.Unsetenv("FOO_TEST_VAR")
	os.Unsetenv("BAZ_TEST_VAR")
	t.Cleanup(func() {
		os.Unsetenv("FOO_TEST_VAR")
		os.Unsetenv("BAZ_TEST_VAR")
	})

	if err := config.LoadEnvFile(envFilePath); err != nil {
		t.Fatalf("unexpected error loading env file: %v", err)
	}

	if value := os.Getenv("FOO_TEST_VAR"); value != "bar" {
		t.Fatalf("expected FOO_TEST_VAR=%q, got %q", "bar", value)
	}

	if value := os.Getenv("BAZ_TEST_VAR"); value != "qux" {
		t.Fatalf("expected BAZ_TEST_VAR=%q, got %q", "qux", value)
	}
}

func TestLoadEnvFile_DoesNotOverrideExistingEnvironmentVariable(t *testing.T) {
	envFilePath := filepath.Join(t.TempDir(), ".env")

	if err := os.WriteFile(envFilePath, []byte("EXISTING_TEST_VAR=fromfile\n"), 0o600); err != nil {
		t.Fatalf("unexpected error writing env file: %v", err)
	}

	os.Setenv("EXISTING_TEST_VAR", "fromenv")
	t.Cleanup(func() {
		os.Unsetenv("EXISTING_TEST_VAR")
	})

	if err := config.LoadEnvFile(envFilePath); err != nil {
		t.Fatalf("unexpected error loading env file: %v", err)
	}

	if value := os.Getenv("EXISTING_TEST_VAR"); value != "fromenv" {
		t.Fatalf("expected existing environment variable to be preserved, got %q", value)
	}
}

func TestLoadEnvFile_ReturnsErrorWhenFileDoesNotExist(t *testing.T) {
	nonExistentPath := filepath.Join(t.TempDir(), "does-not-exist.env")

	if err := config.LoadEnvFile(nonExistentPath); err == nil {
		t.Fatalf("expected an error when the env file does not exist")
	}
}
