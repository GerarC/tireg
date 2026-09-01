package exception

import (
	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/time-registry/domain/util/constant"
)

func NewValidationFailedError(details ...string) *sharedException.DomainError {
	return sharedException.New(constant.CodeValidationFailed, constant.MessageValidationFailed, details...)
}

func NewTimeEntryNotFoundError() *sharedException.DomainError {
	return sharedException.New(constant.CodeTimeEntryNotFound, constant.MessageTimeEntryNotFound)
}
