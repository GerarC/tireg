package entity

import "time"

type AuditEntity struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string    `gorm:"not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string    `gorm:"not null"`
}
