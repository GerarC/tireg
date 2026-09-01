package bootstrap

import (
	"net/http"

	"github.com/gerarc/tireg/internal/health/domain/usecase"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/shared/application/utils/container"
)

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewHealthUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewHealthController)
}

func WireRoutes(mux *http.ServeMux) {
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	healthController := container.MustResolve[*controller.HealthController](appContainer)
	healthController.RegisterRoutes(mux)
}
