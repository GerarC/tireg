package constant

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
)

type ContextKey string

const AuthenticatedUserContextKey ContextKey = "authenticatedUser"
