package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gerarc/tireg/internal/shared/application/utils/logger"
	"github.com/gerarc/tireg/internal/shared/domain/exception"
	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/util/constant"
)

var appLogger logger.Logger

func SetLogger(configuredLogger logger.Logger) {
	appLogger = configuredLogger
}

func WriteError(responseWriter http.ResponseWriter, err error) {
	var domainError *exception.DomainError
	if errors.As(err, &domainError) {
		if appLogger != nil {
			appLogger.Warn(
				"request failed",
				"code", domainError.Code,
				"message", domainError.Message,
				"details", domainError.Details,
			)
		}
	} else {
		if appLogger != nil {
			appLogger.Error("unhandled error", err)
		}
		domainError = exception.New(constant.InternalErrorCode, constant.InternalErrorMessage)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(domainError.Code)
	json.NewEncoder(responseWriter).Encode(ErrorResponseDTO{
		Code:      domainError.Code,
		Message:   domainError.Message,
		Details:   domainError.Details,
		Timestamp: domainError.Timestamp.Format(time.RFC3339),
	})
}

func InvalidRequestBodyError(err error) *exception.DomainError {
	return exception.New(constant.InvalidRequestBodyCode, constant.InvalidRequestBodyMessage, err.Error())
}
