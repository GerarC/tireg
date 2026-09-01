package entity

import (
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/util/constant"
)

type TaskMappingEntity struct {
	ID                   string   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerID              string   `gorm:"type:uuid;not null;index"`
	ProjectLabel         string   `gorm:"not null"`
	Pattern              string   `gorm:"not null"`
	MatchKeywords        []string `gorm:"serializer:json;type:jsonb"`
	MatchOrganizerDomain string
	IssueKey             string
	TypeKey              string
	Notes                string
	sharedEntity.AuditEntity
}

func (TaskMappingEntity) TableName() string {
	return constant.TaskMappingsTableName
}
