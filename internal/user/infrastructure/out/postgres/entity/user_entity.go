package entity

import sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

type UserEntity struct {
	ID           string
	FirstName    string
	LastName     string
	Username     string
	Email        string
	PasswordHash string
	sharedModel.Audit
}
