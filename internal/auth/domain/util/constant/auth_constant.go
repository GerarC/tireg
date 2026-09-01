package constant

const (
	CodeValidationFailed    = 400
	MessageValidationFailed = "AUTH_VALIDATION_FAILED"

	DetailIdentifierRequired = "identifier is required"
	DetailPasswordRequired   = "password is required"

	CodeInvalidCredentials    = 401
	MessageInvalidCredentials = "AUTH_INVALID"
	DetailInvalidCredentials  = "credentials are invalid"
)
