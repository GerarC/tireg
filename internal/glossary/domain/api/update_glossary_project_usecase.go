package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// UpdateGlossaryProjectUseCase exposes the operation to update a glossary project owned by a user.
type UpdateGlossaryProjectUseCase interface {
	// Update validates and updates the glossary project matching the given id, owned by the given user.
	Update(ownerID string, id string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error)
}
