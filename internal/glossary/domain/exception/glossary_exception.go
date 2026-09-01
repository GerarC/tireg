package exception

import (
	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/glossary/domain/util/constant"
)

func NewValidationFailedError(details ...string) *sharedException.DomainError {
	return sharedException.New(constant.CodeValidationFailed, constant.MessageValidationFailed, details...)
}

func NewGlossaryTypeNotFoundError() *sharedException.DomainError {
	return sharedException.New(constant.CodeGlossaryTypeNotFound, constant.MessageGlossaryTypeNotFound)
}

func NewGlossaryProjectNotFoundError() *sharedException.DomainError {
	return sharedException.New(constant.CodeGlossaryProjectNotFound, constant.MessageGlossaryProjectNotFound)
}
