package exception

import (
	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	"github.com/gerarc/tireg/internal/task/domain/util/constant"
)

func NewValidationFailedError(details ...string) *sharedException.DomainError {
	return sharedException.New(constant.CodeValidationFailed, constant.MessageValidationFailed, details...)
}

func NewTaskMappingNotFoundError() *sharedException.DomainError {
	return sharedException.New(constant.CodeTaskMappingNotFound, constant.MessageTaskMappingNotFound)
}
