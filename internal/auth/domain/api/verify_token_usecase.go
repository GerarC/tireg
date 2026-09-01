package api

import "github.com/gerarc/tireg/internal/auth/domain/model"

// VerifyTokenUseCase exposes the operation to authenticate a request from its access token.
type VerifyTokenUseCase interface {
	// Verify validates the given access token and returns the authenticated user it was issued for.
	Verify(token string) (model.AuthenticatedUser, error)
}
