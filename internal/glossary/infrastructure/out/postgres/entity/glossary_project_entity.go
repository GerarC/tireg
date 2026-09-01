package entity

import (
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/util/constant"
)

type GlossaryProjectEntity struct {
	ID             string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerID        string `gorm:"type:uuid;not null;index"`
	ProjectLabel   string `gorm:"not null"`
	Client         string
	JiraProjectKey string
	BoardURL       string
	Notes          string
	sharedEntity.AuditEntity
}

func (GlossaryProjectEntity) TableName() string {
	return constant.GlossaryProjectsTableName
}
