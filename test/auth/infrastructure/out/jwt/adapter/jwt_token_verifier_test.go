package adapter_test

import (
	"testing"
	"time"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"

	"github.com/gerarc/tireg/internal/auth/infrastructure/out/jwt/adapter"
)

func TestVerify_ReturnsAuthenticatedUser_WhenTokenIsValid(t *testing.T) {
	appConfig := &config.Config{JWTSecret: "test-secret", JWTExpiration: time.Hour}

	tokenIssuer := adapter.NewJWTTokenIssuer(appConfig)
	accessToken, err := tokenIssuer.Issue("user-id", "ada")
	if err != nil {
		t.Fatalf("unexpected error issuing token: %v", err)
	}

	tokenVerifier := adapter.NewJWTTokenVerifier(appConfig)

	authenticatedUser, err := tokenVerifier.Verify(accessToken.Token)
	if err != nil {
		t.Fatalf("unexpected error verifying token: %v", err)
	}

	if authenticatedUser.ID != "user-id" || authenticatedUser.Username != "ada" {
		t.Fatalf("unexpected authenticated user: %+v", authenticatedUser)
	}
}

func TestVerify_ReturnsError_WhenTokenIsMalformed(t *testing.T) {
	tokenVerifier := adapter.NewJWTTokenVerifier(&config.Config{JWTSecret: "test-secret"})

	if _, err := tokenVerifier.Verify("not-a-jwt"); err == nil {
		t.Fatalf("expected an error for a malformed token")
	}
}

func TestVerify_ReturnsError_WhenSignedWithDifferentSecret(t *testing.T) {
	issuerConfig := &config.Config{JWTSecret: "issuer-secret", JWTExpiration: time.Hour}
	tokenIssuer := adapter.NewJWTTokenIssuer(issuerConfig)

	accessToken, err := tokenIssuer.Issue("user-id", "ada")
	if err != nil {
		t.Fatalf("unexpected error issuing token: %v", err)
	}

	tokenVerifier := adapter.NewJWTTokenVerifier(&config.Config{JWTSecret: "different-secret"})

	if _, err := tokenVerifier.Verify(accessToken.Token); err == nil {
		t.Fatalf("expected an error when the token was signed with a different secret")
	}
}

func TestVerify_ReturnsError_WhenTokenExpired(t *testing.T) {
	appConfig := &config.Config{JWTSecret: "test-secret", JWTExpiration: -time.Hour}
	tokenIssuer := adapter.NewJWTTokenIssuer(appConfig)

	accessToken, err := tokenIssuer.Issue("user-id", "ada")
	if err != nil {
		t.Fatalf("unexpected error issuing token: %v", err)
	}

	tokenVerifier := adapter.NewJWTTokenVerifier(appConfig)

	if _, err := tokenVerifier.Verify(accessToken.Token); err == nil {
		t.Fatalf("expected an error for an expired token")
	}
}
