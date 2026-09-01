package spi

import "github.com/gerarc/tireg/internal/auth/domain/model"

// TokenIssuer defines the contract to issue signed access tokens for an authenticated user.
type TokenIssuer interface {
	// Issue generates a signed access token for the given user, returning the token and its expiration time.
	Issue(userID string, username string) (model.AccessToken, error)
}
