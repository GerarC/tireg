package logger

// Logger defines the contract for structured application logging, independent of the underlying logging library.
type Logger interface {
	// Info logs a message at info level with the given structured key/value attributes.
	Info(msg string, args ...any)
	// Warn logs a message at warn level with the given structured key/value attributes.
	Warn(msg string, args ...any)
	// Error logs a message at error level, attaching the given error and structured key/value attributes.
	Error(msg string, err error, args ...any)
	// Debug logs a message at debug level with the given structured key/value attributes.
	Debug(msg string, args ...any)
	// With returns a Logger that always includes the given structured key/value attributes.
	With(args ...any) Logger
}
