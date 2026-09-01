package adapter_test

import (
	"testing"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"

	"github.com/gerarc/tireg/internal/auth/infrastructure/out/jwt/adapter"
)

func TestIssue_ReturnsTokenWithExpectedClaims(t *testing.T) {
	tokenIssuer := adapter.NewJWTTokenIssuer(&config.Config{JWTSecret: "test-secret", JWTExpiration: time.Hour})

	accessToken, err := tokenIssuer.Issue("user-id", "ada")
	if err != nil {
		t.Fatalf("unexpected error issuing token: %v", err)
	}

	parsedToken, err := jwtlib.Parse(accessToken.Token, func(token *jwtlib.Token) (any, error) {
		return []byte("test-secret"), nil
	})
	if err != nil {
		t.Fatalf("unexpected error parsing token: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwtlib.MapClaims)
	if !ok || !parsedToken.Valid {
		t.Fatalf("expected valid claims, got %v", parsedToken.Claims)
	}

	if claims["sub"] != "user-id" {
		t.Fatalf("expected subject %q, got %v", "user-id", claims["sub"])
	}

	if claims["username"] != "ada" {
		t.Fatalf("expected username %q, got %v", "ada", claims["username"])
	}
}

func TestIssue_SetsExpirationAccordingToConfig(t *testing.T) {
	tokenIssuer := adapter.NewJWTTokenIssuer(&config.Config{JWTSecret: "test-secret", JWTExpiration: 30 * time.Minute})

	before := time.Now()
	accessToken, err := tokenIssuer.Issue("user-id", "ada")
	after := time.Now()

	if err != nil {
		t.Fatalf("unexpected error issuing token: %v", err)
	}

	minExpected := before.Add(30 * time.Minute)
	maxExpected := after.Add(30 * time.Minute)

	if accessToken.ExpiresAt.Before(minExpected) || accessToken.ExpiresAt.After(maxExpected) {
		t.Fatalf("expected expiration between %v and %v, got %v", minExpected, maxExpected, accessToken.ExpiresAt)
	}
}
