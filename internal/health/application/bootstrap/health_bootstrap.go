package bootstrap

import (
	"net/http"

	"github.com/gerarc/tireg/internal/health/domain/usecase"
	"github.com/gerarc/tireg/internal/health/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/shared/application/utils/container"
)

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewCheckHealthUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewCheckHealthController)
}

func WireRoutes(mux *http.ServeMux) {
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	checkHealthController := container.MustResolve[*controller.CheckHealthController](appContainer)
	checkHealthController.RegisterRoutes(mux)
}
