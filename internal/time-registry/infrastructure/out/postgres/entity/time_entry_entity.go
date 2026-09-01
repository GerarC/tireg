package entity

import (
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/util/constant"
)

type TimeEntryEntity struct {
	ID            string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerID       string `gorm:"type:uuid;not null;index:time_entries_owner_id_date_idx"`
	Date          string `gorm:"not null;index:time_entries_owner_id_date_idx"`
	ProjectLabel  string `gorm:"not null"`
	TypeKey       string
	IssueKey      string
	Start         string  `gorm:"not null"`
	End           string  `gorm:"not null"`
	Hours         float64 `gorm:"not null"`
	Description   string
	JiraWorklogID string
	sharedEntity.AuditEntity
}

func (TimeEntryEntity) TableName() string {
	return constant.TimeEntriesTableName
}
