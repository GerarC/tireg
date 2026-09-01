package spi

import "github.com/gerarc/tireg/internal/user/domain/model"

// UserQueryRepository defines the read persistence operations required by the user use cases.
type UserQueryRepository interface {
	// ExistsByUsername reports whether a user with the given username already exists.
	ExistsByUsername(username string) (bool, error)
	// ExistsByEmail reports whether a user with the given email already exists.
	ExistsByEmail(email string) (bool, error)
	// FindByIdentifier returns the user whose username or email matches the given identifier.
	FindByIdentifier(identifier string) (model.User, error)
}
