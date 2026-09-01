package bootstrap

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/shared/application/utils/container"

	"github.com/gerarc/tireg/internal/task/domain/usecase"
	"github.com/gerarc/tireg/internal/task/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/adapter"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/entity"
	"github.com/gerarc/tireg/internal/task/infrastructure/out/postgres/repository"
)

func WirePersistenceDependencies() {
	appContainer := container.GetInstance()

	db := container.MustResolve[*gorm.DB](appContainer)
	if err := db.AutoMigrate(&entity.TaskMappingEntity{}); err != nil {
		panic(err)
	}

	appContainer.Register(repository.NewTaskMappingCommandRepository)
	appContainer.Register(repository.NewTaskMappingQueryRepository)
	appContainer.Register(adapter.NewTaskMappingCommandAdapter)
	appContainer.Register(adapter.NewTaskMappingQueryAdapter)
}

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewCreateTaskMappingUseCase)
	appContainer.Register(usecase.NewListTaskMappingsUseCase)
	appContainer.Register(usecase.NewFindTaskMappingByIDUseCase)
	appContainer.Register(usecase.NewUpdateTaskMappingUseCase)
	appContainer.Register(usecase.NewDeleteTaskMappingUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewCreateTaskMappingController)
	appContainer.Register(controller.NewListTaskMappingsController)
	appContainer.Register(controller.NewFindTaskMappingByIDController)
	appContainer.Register(controller.NewUpdateTaskMappingController)
	appContainer.Register(controller.NewDeleteTaskMappingController)
}

func WireRoutes(mux *http.ServeMux) {
	WirePersistenceDependencies()
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	container.MustResolve[*controller.CreateTaskMappingController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.ListTaskMappingsController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.FindTaskMappingByIDController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.UpdateTaskMappingController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.DeleteTaskMappingController](appContainer).RegisterRoutes(mux)
}
