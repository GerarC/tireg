package spi

import "github.com/gerarc/tireg/internal/auth/domain/model"

// TokenVerifier defines the contract to validate a signed access token and extract the authenticated user from it.
type TokenVerifier interface {
	// Verify validates the given token and returns the authenticated user it was issued for.
	Verify(token string) (model.AuthenticatedUser, error)
}
