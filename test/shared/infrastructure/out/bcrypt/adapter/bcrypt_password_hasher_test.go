package adapter_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/shared/infrastructure/out/bcrypt/adapter"
)

func TestHash_ProducesDifferentHashesForSamePassword(t *testing.T) {
	passwordHasher := adapter.NewBcryptPasswordHasher()

	firstHash, err := passwordHasher.Hash("super-secret-password")
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	secondHash, err := passwordHasher.Hash("super-secret-password")
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if firstHash == secondHash {
		t.Fatalf("expected different hashes for the same password due to salting")
	}
}

func TestVerify_ReturnsTrue_WhenPasswordMatches(t *testing.T) {
	passwordHasher := adapter.NewBcryptPasswordHasher()

	hash, err := passwordHasher.Hash("super-secret-password")
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if !passwordHasher.Verify(hash, "super-secret-password") {
		t.Fatalf("expected password to match its own hash")
	}
}

func TestVerify_ReturnsFalse_WhenPasswordDoesNotMatch(t *testing.T) {
	passwordHasher := adapter.NewBcryptPasswordHasher()

	hash, err := passwordHasher.Hash("super-secret-password")
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if passwordHasher.Verify(hash, "wrong-password") {
		t.Fatalf("expected wrong password not to match the hash")
	}
}
