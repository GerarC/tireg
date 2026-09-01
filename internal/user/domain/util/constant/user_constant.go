package constant

const (
	CodeAlreadyTaken    = 409
	MessageAlreadyTaken = "USER_ALREADY_TAKEN"

	CodeValidationFailed    = 400
	MessageValidationFailed = "USER_VALIDATION_FAILED"

	CodeUserNotFound    = 404
	MessageUserNotFound = "USER_NOT_FOUND"

	DetailUsernameAlreadyTaken = "username already taken"
	DetailEmailAlreadyTaken    = "email already registered"

	DetailFirstNameRequired = "first name is required"
	DetailLastNameRequired  = "last name is required"
	DetailUsernameTooShort  = "username must be at least 3 characters"
	DetailPasswordTooShort  = "password must be at least 8 characters"
	DetailInvalidEmail      = "email must be a valid email address"

	MinUsernameLength = 3
	MinPasswordLength = 8

	SelfRegistrationActor = "self-registration"
)
