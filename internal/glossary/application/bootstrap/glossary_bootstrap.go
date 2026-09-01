package bootstrap

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gerarc/tireg/internal/shared/application/utils/container"

	"github.com/gerarc/tireg/internal/glossary/domain/usecase"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/in/rest/v1/controller"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/adapter"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/entity"
	"github.com/gerarc/tireg/internal/glossary/infrastructure/out/postgres/repository"
)

func WirePersistenceDependencies() {
	appContainer := container.GetInstance()

	db := container.MustResolve[*gorm.DB](appContainer)
	if err := db.AutoMigrate(&entity.GlossaryTypeEntity{}, &entity.GlossaryProjectEntity{}); err != nil {
		panic(err)
	}

	appContainer.Register(repository.NewGlossaryTypeCommandRepository)
	appContainer.Register(repository.NewGlossaryTypeQueryRepository)
	appContainer.Register(repository.NewGlossaryProjectCommandRepository)
	appContainer.Register(repository.NewGlossaryProjectQueryRepository)

	appContainer.Register(adapter.NewGlossaryTypeCommandAdapter)
	appContainer.Register(adapter.NewGlossaryTypeQueryAdapter)
	appContainer.Register(adapter.NewGlossaryProjectCommandAdapter)
	appContainer.Register(adapter.NewGlossaryProjectQueryAdapter)
}

func WireUseCaseDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(usecase.NewCreateGlossaryTypeUseCase)
	appContainer.Register(usecase.NewListGlossaryTypesUseCase)
	appContainer.Register(usecase.NewUpdateGlossaryTypeUseCase)
	appContainer.Register(usecase.NewDeleteGlossaryTypeUseCase)

	appContainer.Register(usecase.NewCreateGlossaryProjectUseCase)
	appContainer.Register(usecase.NewListGlossaryProjectsUseCase)
	appContainer.Register(usecase.NewUpdateGlossaryProjectUseCase)
	appContainer.Register(usecase.NewDeleteGlossaryProjectUseCase)

	appContainer.Register(usecase.NewGetGlossaryUseCase)
}

func WireHandlerDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(controller.NewGetGlossaryController)

	appContainer.Register(controller.NewCreateGlossaryTypeController)
	appContainer.Register(controller.NewListGlossaryTypesController)
	appContainer.Register(controller.NewUpdateGlossaryTypeController)
	appContainer.Register(controller.NewDeleteGlossaryTypeController)

	appContainer.Register(controller.NewCreateGlossaryProjectController)
	appContainer.Register(controller.NewListGlossaryProjectsController)
	appContainer.Register(controller.NewUpdateGlossaryProjectController)
	appContainer.Register(controller.NewDeleteGlossaryProjectController)
}

func WireRoutes(mux *http.ServeMux) {
	WirePersistenceDependencies()
	WireUseCaseDependencies()
	WireHandlerDependencies()

	appContainer := container.GetInstance()

	container.MustResolve[*controller.GetGlossaryController](appContainer).RegisterRoutes(mux)

	container.MustResolve[*controller.CreateGlossaryTypeController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.ListGlossaryTypesController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.UpdateGlossaryTypeController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.DeleteGlossaryTypeController](appContainer).RegisterRoutes(mux)

	container.MustResolve[*controller.CreateGlossaryProjectController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.ListGlossaryProjectsController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.UpdateGlossaryProjectController](appContainer).RegisterRoutes(mux)
	container.MustResolve[*controller.DeleteGlossaryProjectController](appContainer).RegisterRoutes(mux)
}
