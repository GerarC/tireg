package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// CreateGlossaryTypeUseCase exposes the operation to create a new glossary type owned by a user.
type CreateGlossaryTypeUseCase interface {
	// Create validates and persists a new glossary type for the given owner.
	Create(ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error)
}
