package spi

import "github.com/gerarc/tireg/internal/glossary/domain/model"

// GlossaryProjectQueryRepository defines the read persistence operations required by the glossary project use cases.
type GlossaryProjectQueryRepository interface {
	// SelectAllByOwner returns every glossary project owned by the given user.
	SelectAllByOwner(ownerID string) ([]model.GlossaryProject, error)
}
