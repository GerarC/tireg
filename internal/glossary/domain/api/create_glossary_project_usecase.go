package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// CreateGlossaryProjectUseCase exposes the operation to create a new glossary project owned by a user.
type CreateGlossaryProjectUseCase interface {
	// Create validates and persists a new glossary project for the given owner.
	Create(ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error)
}
