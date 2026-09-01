package entity

import (
	sharedEntity "github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres/entity"

	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/util/constant"
)

type UserEntity struct {
	ID           string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Username     string `gorm:"uniqueIndex:users_username_key;not null"`
	Email        string `gorm:"uniqueIndex:users_email_key;not null"`
	PasswordHash string `gorm:"not null"`
	sharedEntity.AuditEntity
}

func (UserEntity) TableName() string {
	return constant.UsersTableName
}
