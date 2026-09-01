package api

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// GetGlossaryUseCase exposes the operation to fetch the whole glossary (types and projects) owned by a user.
type GetGlossaryUseCase interface {
	// Get returns the full glossary owned by the given user.
	Get(ownerID string) (model.Glossary, error)
}
