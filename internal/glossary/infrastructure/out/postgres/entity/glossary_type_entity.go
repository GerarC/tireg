package entity

import (
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/util/constant"
)

type GlossaryTypeEntity struct {
	ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerID     string `gorm:"type:uuid;not null;uniqueIndex:glossary_types_owner_id_type_key_key"`
	TypeKey     string `gorm:"not null;uniqueIndex:glossary_types_owner_id_type_key_key"`
	Label       string `gorm:"not null"`
	Description string
	sharedEntity.AuditEntity
}

func (GlossaryTypeEntity) TableName() string {
	return constant.GlossaryTypesTableName
}
