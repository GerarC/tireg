package api

import "github.com/gerarc/tireg/internal/user/domain/model"

// FindUserByIdentifierUseCase exposes the operation to find a user by username or email.
type FindUserByIdentifierUseCase interface {
	// FindByIdentifier returns the user whose username or email matches the given identifier.
	FindByIdentifier(identifier string) (model.User, error)
}
