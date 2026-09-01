package middleware

import (
	"context"
	"net/http"
	"strings"

	sharedRest "github.com/gerarc/tireg/internal/shared/infrastructure/in/rest"
	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware/util/constant"

	authApi "github.com/gerarc/tireg/internal/auth/domain/api"
	authModel "github.com/gerarc/tireg/internal/auth/domain/model"
)

type RequireAuthMiddleware struct {
	verifyTokenUseCase authApi.VerifyTokenUseCase
}

func NewRequireAuthMiddleware(verifyTokenUseCase authApi.VerifyTokenUseCase) *RequireAuthMiddleware {
	return &RequireAuthMiddleware{verifyTokenUseCase: verifyTokenUseCase}
}

func (requireAuthMiddleware *RequireAuthMiddleware) Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		token := strings.TrimPrefix(request.Header.Get(constant.AuthorizationHeader), constant.BearerPrefix)

		authenticatedUser, err := requireAuthMiddleware.verifyTokenUseCase.Verify(token)
		if err != nil {
			sharedRest.WriteError(responseWriter, err)
			return
		}

		ctx := context.WithValue(request.Context(), constant.AuthenticatedUserContextKey, authenticatedUser)
		next(responseWriter, request.WithContext(ctx))
	}
}

func AuthenticatedUserFromContext(ctx context.Context) (authModel.AuthenticatedUser, bool) {
	authenticatedUser, ok := ctx.Value(constant.AuthenticatedUserContextKey).(authModel.AuthenticatedUser)
	return authenticatedUser, ok
}
