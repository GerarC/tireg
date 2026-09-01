package adapter

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/gerarc/tireg/internal/shared/domain/spi"
)

type BcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() spi.PasswordHasher {
	return &BcryptPasswordHasher{}
}

func (bcryptPasswordHasher *BcryptPasswordHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (bcryptPasswordHasher *BcryptPasswordHasher) Verify(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
