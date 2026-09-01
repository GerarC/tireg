package model

import sharedModel "github.com/gerarc/tireg/internal/shared/domain/model"

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Username     string
	Email        string
	PasswordHash string
	sharedModel.Audit
}
