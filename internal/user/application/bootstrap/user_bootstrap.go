package bootstrap

import (
	"net/http"

	"github.com/gerarc/tireg/internal/shared/application/utils/container"

	"github.com/gerarc/tireg/internal/user/domain/usecase"
	"github.com/gerarc/tireg/internal/user/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/adapter"
	"github.com/gerarc/tireg/internal/user/infrastructure/out/postgres/repository"
)

func WirePersistenceDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(repository.NewUserCommandRepository)
	appContainer.Register(repository.NewUserQueryRepository)
	appContainer.Register(adapter.NewUserCommandAdapter)
	appContainer.Register(adapter.NewUserQueryAdapter)
}

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewRegisterUserUseCase)
	appContainer.Register(usecase.NewFindUserByIdentifierUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewRegisterUserController)
}

func WireRoutes(mux *http.ServeMux) {
	WirePersistenceDependencies()
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	registerUserController := container.MustResolve[*controller.RegisterUserController](appContainer)
	registerUserController.RegisterRoutes(mux)
}
