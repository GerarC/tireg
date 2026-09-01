package api

import "github.com/gerarc/tireg/internal/auth/domain/model"

// LoginUseCase exposes the operation to authenticate a user.
type LoginUseCase interface {
	// Login validates the given credentials and returns a signed access token.
	Login(credentials model.Credentials) (model.AccessToken, error)
}
