package constant

const (
	CodeValidationFailed    = 400
	MessageValidationFailed = "TIME_ENTRY_VALIDATION_FAILED"

	CodeTimeEntryNotFound    = 404
	MessageTimeEntryNotFound = "TIME_ENTRY_NOT_FOUND"

	DetailDateRequired         = "date is required"
	DetailProjectLabelRequired = "project label is required"
	DetailStartRequired        = "start is required"
	DetailEndRequired          = "end is required"
	DetailHoursMustBePositive  = "hours must be greater than zero"
)
