package constant

import "net/http"

const (
	InternalErrorCode    = http.StatusInternalServerError
	InternalErrorMessage = "INTERNAL_ERROR"

	InvalidRequestBodyCode    = http.StatusBadRequest
	InvalidRequestBodyMessage = "INVALID_REQUEST_BODY"
)
