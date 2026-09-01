package exception

import (
	"strings"
	"time"
)

type DomainError struct {
	Code      int
	Message   string
	Details   []string
	Timestamp time.Time
}

func New(code int, message string, details ...string) *DomainError {
	return &DomainError{
		Code:      code,
		Message:   message,
		Details:   details,
		Timestamp: time.Now(),
	}
}

func (domainError *DomainError) Error() string {
	if len(domainError.Details) == 0 {
		return domainError.Message
	}

	return domainError.Message + ": " + strings.Join(domainError.Details, "; ")
}
