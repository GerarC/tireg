package spi

import "github.com/gerarc/tireg/internal/user/domain/model"

// UserCommandRepository defines the write persistence operations required by the user use cases.
type UserCommandRepository interface {
	// Save persists a new user and returns it with its generated ID and audit timestamps.
	Save(user model.User) (model.User, error)
}
