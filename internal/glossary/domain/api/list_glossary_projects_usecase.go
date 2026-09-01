package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// ListGlossaryProjectsUseCase exposes the operation to list every glossary project owned by a user.
type ListGlossaryProjectsUseCase interface {
	// List returns every glossary project owned by the given user.
	List(ownerID string) ([]model.GlossaryProject, error)
}
