package spi

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// GlossaryProjectCommandRepository defines the write persistence operations required by the glossary project use cases.
type GlossaryProjectCommandRepository interface {
	// Insert persists a new glossary project and returns it with its generated ID and audit timestamps.
	Insert(glossaryProject model.GlossaryProject) (model.GlossaryProject, error)
	// UpdateByIDAndOwner updates the glossary project matching the given id and owner, returning it updated.
	UpdateByIDAndOwner(id string, ownerID string, glossaryProject model.GlossaryProject) (model.GlossaryProject, error)
	// DeleteByIDAndOwner deletes the glossary project matching the given id and owner.
	DeleteByIDAndOwner(id string, ownerID string) error
}
