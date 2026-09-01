package constant

const (
	EnvFileName = ".env"
	PortEnvVar  = "PORT"
	DefaultPort = "8080"

	PostgresHostEnvVar     = "POSTGRES_HOST"
	PostgresPortEnvVar     = "POSTGRES_PORT"
	PostgresUserEnvVar     = "POSTGRES_USER"
	PostgresPasswordEnvVar = "POSTGRES_PASSWORD"
	PostgresDBNameEnvVar   = "POSTGRES_DB"
	PostgresSSLModeEnvVar  = "POSTGRES_SSLMODE"

	DefaultPostgresHost    = "localhost"
	DefaultPostgresPort    = "5432"
	DefaultPostgresUser    = "postgres"
	DefaultPostgresDBName  = "tireg"
	DefaultPostgresSSLMode = "disable"

	JWTSecretEnvVar            = "JWT_SECRET"
	JWTExpirationMinutesEnvVar = "JWT_EXPIRATION_MINUTES"

	DefaultJWTExpirationMinutes = "60"
)
