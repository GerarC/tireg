package spi

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// GlossaryTypeQueryRepository defines the read persistence operations required by the glossary type use cases.
type GlossaryTypeQueryRepository interface {
	// SelectAllByOwner returns every glossary type owned by the given user.
	SelectAllByOwner(ownerID string) ([]model.GlossaryType, error)
}
