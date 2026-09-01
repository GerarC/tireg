package exception

import (
	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/user/domain/util/constant"
)

func NewAlreadyTakenError(details ...string) *sharedException.DomainError {
	return sharedException.New(constant.CodeAlreadyTaken, constant.MessageAlreadyTaken, details...)
}

func NewValidationFailedError(details ...string) *sharedException.DomainError {
	return sharedException.New(constant.CodeValidationFailed, constant.MessageValidationFailed, details...)
}

func NewUserNotFoundError() *sharedException.DomainError {
	return sharedException.New(constant.CodeUserNotFound, constant.MessageUserNotFound)
}
