package bootstrap

import (
	"net/http"

	"github.com/gerarc/tireg/internal/shared/application/utils/container"

	jwtadapter "github.com/gerarc/tireg/internal/auth/infrastructure/out/jwt/adapter"

	"github.com/gerarc/tireg/internal/auth/domain/usecase"
	"github.com/gerarc/tireg/internal/auth/infrastructure/in/rest/v1/controller"
)

func WireTokenIssuerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(jwtadapter.NewJWTTokenIssuer)
	appContainer.Register(jwtadapter.NewJWTTokenVerifier)
}

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewLoginUseCase)
	appContainer.Register(usecase.NewVerifyTokenUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewLoginController)
}

func WireRoutes(mux *http.ServeMux) {
	WireTokenIssuerDependencies()
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	loginController := container.MustResolve[*controller.LoginController](appContainer)
	loginController.RegisterRoutes(mux)
}
