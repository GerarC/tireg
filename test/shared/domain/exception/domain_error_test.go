package exception_test

import (
	"testing"
	"time"

	"github.com/gerarc/tireg/internal/shared/domain/exception"
)

func TestNew_SetsAllFieldsAndTimestamp(t *testing.T) {
	before := time.Now()

	domainError := exception.New(409, "USER_ALREADY_TAKEN", "username already taken", "email already registered")

	after := time.Now()

	if domainError.Code != 409 {
		t.Fatalf("expected code 409, got %d", domainError.Code)
	}

	if domainError.Message != "USER_ALREADY_TAKEN" {
		t.Fatalf("expected message %q, got %q", "USER_ALREADY_TAKEN", domainError.Message)
	}

	if len(domainError.Details) != 2 {
		t.Fatalf("expected 2 details, got %d", len(domainError.Details))
	}

	if domainError.Timestamp.Before(before) || domainError.Timestamp.After(after) {
		t.Fatalf("expected timestamp between %v and %v, got %v", before, after, domainError.Timestamp)
	}
}

func TestError_IncludesDetailsWhenPresent(t *testing.T) {
	domainError := exception.New(400, "USER_VALIDATION_FAILED", "username must be at least 3 characters", "password must be at least 8 characters")

	errorMessage := domainError.Error()

	if errorMessage != "USER_VALIDATION_FAILED: username must be at least 3 characters; password must be at least 8 characters" {
		t.Fatalf("unexpected error message: %q", errorMessage)
	}
}

func TestError_OmitsDetailsWhenAbsent(t *testing.T) {
	domainError := exception.New(404, "USER_NOT_FOUND")

	if domainError.Error() != "USER_NOT_FOUND" {
		t.Fatalf("expected error message %q, got %q", "USER_NOT_FOUND", domainError.Error())
	}
}
