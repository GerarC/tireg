package api

import "github.com/gerarc/tireg/internal/user/domain/model"

// RegisterUserUseCase exposes the operation to register a new user.
type RegisterUserUseCase interface {
	// Register creates a new user with a hashed password, after validating its fields and that the username and email are not already taken.
	Register(registration model.UserRegistration) (model.User, error)
}
