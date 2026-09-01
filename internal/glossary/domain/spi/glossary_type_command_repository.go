package spi

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// GlossaryTypeCommandRepository defines the write persistence operations required by the glossary type use cases.
type GlossaryTypeCommandRepository interface {
	// Insert persists a new glossary type and returns it with its generated ID and audit timestamps.
	Insert(glossaryType model.GlossaryType) (model.GlossaryType, error)
	// UpdateByIDAndOwner updates the glossary type matching the given id and owner, returning it updated.
	UpdateByIDAndOwner(id string, ownerID string, glossaryType model.GlossaryType) (model.GlossaryType, error)
	// DeleteByIDAndOwner deletes the glossary type matching the given id and owner.
	DeleteByIDAndOwner(id string, ownerID string) error
}
