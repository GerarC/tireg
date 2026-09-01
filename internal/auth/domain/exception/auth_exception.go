package exception

import (
	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/auth/domain/util/constant"
)

func NewValidationFailedError(details ...string) *sharedException.DomainError {
	return sharedException.New(constant.CodeValidationFailed, constant.MessageValidationFailed, details...)
}

func NewInvalidCredentialsError() *sharedException.DomainError {
	return sharedException.New(constant.CodeInvalidCredentials, constant.MessageInvalidCredentials, constant.DetailInvalidCredentials)
}
