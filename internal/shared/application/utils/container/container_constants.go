package container

const (
	errMustBeFunction         = "constructor must be a function, got %v"
	errInvalidOutputs         = "constructor of '%s' must return one or two elements"
	errSecondMustBeError      = "constructor '%s': second returned element must be an error"
	errNoProvider             = "no provider registered for name %s"
	errResolveFailed          = "failed to resolve dependency for provider %s: %w"
	errConstructorError       = "constructor for '%s' returned an error: %w"
	errDependencyInjection    = "failed resolving dependencies for %s: %w"
	errComponentNotRegistered = "component of type %s not registered in container"
	errComponentNotResolved   = "failed to resolve component %s: %v"
	errFailedToAutoRegister   = "failed to %s (Type: %s): %w"
	errConstructorTypeDup     = "contructor for type '%s' already registered"

	errorTypeName = "error"
)

const (
	returnedInstanceIndex            = 0
	returnedErrorIndex               = 1
	emptyReturnedList                = 0
	constructorReturnedElementsLimit = 2
)

const (
	transientKeyword = "Transient"
)

const (
	empty = 0
)

const (
	SQS_CLIENT = "SQS_CLIENT"
)
