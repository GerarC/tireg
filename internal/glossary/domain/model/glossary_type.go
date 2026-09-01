package model

import sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

type GlossaryType struct {
	ID          string
	OwnerID     string
	TypeKey     string
	Label       string
	Description string
	sharedModel.Audit
}
