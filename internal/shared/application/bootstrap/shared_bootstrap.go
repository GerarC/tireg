package bootstrap

import (
	"github.com/gerarc/tireg/internal/shared/infrastructure/in/rest/middleware"
	bcryptadapter "github.com/gerarc/tireg/internal/shared/infrastructure/out/bcrypt/adapter"
	"github.com/gerarc/tireg/internal/shared/infrastructure/out/postgres"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/application/utils/container"
)

func WireConfigDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(config.NewConfig)
}

func WirePostgresDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(postgres.GetClient)
}

func WirePasswordHasherDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(bcryptadapter.NewBcryptPasswordHasher)
}

func WireMiddlewareDependencies() {
	appContainer := container.GetInstance()

	appContainer.Register(middleware.NewRequireAuthMiddleware)
}

func WireInfrastructureDependencies() {
	WireConfigDependencies()
	WirePostgresDependencies()
	WirePasswordHasherDependencies()
	WireMiddlewareDependencies()
}
