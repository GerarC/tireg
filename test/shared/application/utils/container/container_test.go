package container_test

import (
	"testing"

	"github.com/gerarc/tireg/internal/shared/application/utils/container"
)

type Greeter interface {
	Greet() string
}

type greeterImplemented struct {
	message string
}

func (greeter *greeterImplemented) Greet() string {
	return greeter.message
}

func NewGreeterSingleton() Greeter {
	return &greeterImplemented{message: "singleton"}
}

func NewGreeterTransient() Greeter {
	return &greeterImplemented{message: "transient"}
}

type Repository interface {
	Name() string
}

type repositoryImplemented struct {
	name string
}

func (repository *repositoryImplemented) Name() string {
	return repository.name
}

func NewPrimaryRepository() Repository {
	return &repositoryImplemented{name: "primary"}
}

func NewSecondaryRepository() Repository {
	return &repositoryImplemented{name: "secondary"}
}

type ServiceParams struct {
	Repository Repository `inject:"primary"`
}

type Service interface {
	RepositoryName() string
}

type serviceImplemented struct {
	repository Repository
}

func (service *serviceImplemented) RepositoryName() string {
	return service.repository.Name()
}

func NewService(params ServiceParams) Service {
	return &serviceImplemented{repository: params.Repository}
}

func TestMustResolve_ReturnsSameInstanceForSingleton(t *testing.T) {
	appContainer := container.NewContainer()

	if err := appContainer.Register(NewGreeterSingleton); err != nil {
		t.Fatalf("unexpected error registering constructor: %v", err)
	}

	first := container.MustResolve[Greeter](appContainer)
	second := container.MustResolve[Greeter](appContainer)

	if first != second {
		t.Fatalf("expected singleton to return the same instance on every resolve")
	}
}

func TestMustResolve_ReturnsNewInstanceForTransient(t *testing.T) {
	appContainer := container.NewContainer()

	if err := appContainer.Register(NewGreeterTransient); err != nil {
		t.Fatalf("unexpected error registering constructor: %v", err)
	}

	first := container.MustResolve[Greeter](appContainer)
	second := container.MustResolve[Greeter](appContainer)

	if first == second {
		t.Fatalf("expected transient to return a new instance on every resolve")
	}
}

func TestMustResolve_InjectsNamedDependencyIntoStructParam(t *testing.T) {
	appContainer := container.NewContainer()

	if err := appContainer.Register(NewPrimaryRepository, "primary"); err != nil {
		t.Fatalf("unexpected error registering primary repository: %v", err)
	}
	if err := appContainer.Register(NewSecondaryRepository, "secondary"); err != nil {
		t.Fatalf("unexpected error registering secondary repository: %v", err)
	}
	if err := appContainer.Register(NewService); err != nil {
		t.Fatalf("unexpected error registering service: %v", err)
	}

	service := container.MustResolve[Service](appContainer)

	if service.RepositoryName() != "primary" {
		t.Fatalf("expected injected repository %q, got %q", "primary", service.RepositoryName())
	}
}

func TestMustResolveNamed_ResolvesSpecificImplementation(t *testing.T) {
	appContainer := container.NewContainer()

	if err := appContainer.Register(NewPrimaryRepository, "primary"); err != nil {
		t.Fatalf("unexpected error registering primary repository: %v", err)
	}
	if err := appContainer.Register(NewSecondaryRepository, "secondary"); err != nil {
		t.Fatalf("unexpected error registering secondary repository: %v", err)
	}

	secondary := container.MustResolveNamed[Repository](appContainer, "secondary")

	if secondary.Name() != "secondary" {
		t.Fatalf("expected resolved repository %q, got %q", "secondary", secondary.Name())
	}
}

func TestMustResolve_PanicsWhenTypeNotRegistered(t *testing.T) {
	appContainer := container.NewContainer()

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatalf("expected panic when resolving an unregistered type")
		}
	}()

	container.MustResolve[Greeter](appContainer)
}
