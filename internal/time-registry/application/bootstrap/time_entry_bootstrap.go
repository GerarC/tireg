package bootstrap

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/shared/application/utils/container"

	"github.com/gerarc/tireg/internal/time-registry/domain/usecase"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/adapter"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/entity"
	"github.com/gerarc/tireg/internal/time-registry/infrastructure/out/postgres/repository"
)

func WirePersistenceDependencies() {
	appContainer := container.GetInstance()

	db := container.MustResolve[*gorm.DB](appContainer)
	if err := db.AutoMigrate(&entity.TimeEntryEntity{}); err != nil {
		panic(err)
	}

	appContainer.Register(repository.NewTimeEntryCommandRepository)
	appContainer.Register(repository.NewTimeEntryQueryRepository)
	appContainer.Register(adapter.NewTimeEntryCommandAdapter)
	appContainer.Register(adapter.NewTimeEntryQueryAdapter)
}

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewCreateTimeEntryUseCase)
	appContainer.Register(usecase.NewListTimeEntriesUseCase)
	appContainer.Register(usecase.NewFindTimeEntryByIDUseCase)
	appContainer.Register(usecase.NewUpdateTimeEntryUseCase)
	appContainer.Register(usecase.NewDeleteTimeEntryUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewCreateTimeEntryController)
	appContainer.Register(controller.NewListTimeEntriesController)
	appContainer.Register(controller.NewFindTimeEntryByIDController)
	appContainer.Register(controller.NewUpdateTimeEntryController)
	appContainer.Register(controller.NewDeleteTimeEntryController)
}

func WireRoutes(mux *http.ServeMux) {
	WirePersistenceDependencies()
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	container.MustResolve[*controller.CreateTimeEntryController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.ListTimeEntriesController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.FindTimeEntryByIDController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.UpdateTimeEntryController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.DeleteTimeEntryController](appContainer).RegisterRoutes(mux)
}
