package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// UpdateGlossaryTypeUseCase exposes the operation to update a glossary type owned by a user.
type UpdateGlossaryTypeUseCase interface {
	// Update validates and updates the glossary type matching the given id, owned by the given user.
	Update(ownerID string, id string, glossaryType model.GlossaryType) (model.GlossaryType, error)
}
