package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// ListGlossaryTypesUseCase exposes the operation to list every glossary type owned by a user.
type ListGlossaryTypesUseCase interface {
	// List returns every glossary type owned by the given user, seeding the default types on first access.
	List(ownerID string) ([]model.GlossaryType, error)
}
