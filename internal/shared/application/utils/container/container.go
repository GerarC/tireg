package container

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Container struct {
	providers providerMap
}

func NewContainer() *Container {
	providers := make(providerMap)
	return &Container{
		providers: providers,
	}
}

func (container *Container) Register(constructor any, optionalName ...string) error {
	originalConstructor := reflect.ValueOf(constructor)
	constructorType := originalConstructor.Type()

	registrationName := ""
	if len(optionalName) > 0 {
		registrationName = optionalName[0]
	}

	funcName := runtime.FuncForPC(originalConstructor.Pointer()).Name()
	parts := strings.Split(funcName, ".")
	simpleFuncName := parts[len(parts)-1]

	pType := SINGLETON
	if strings.Contains(simpleFuncName, "Transient") {
		pType = TRANSIENT
	} else if strings.Contains(simpleFuncName, "Request") {
		pType = REQUEST
	}

	serviceType := constructorType.Out(returnedInstanceIndex)

	if err := container.registerComponent(constructor, pType, registrationName); err != nil {
		return fmt.Errorf(errFailedToAutoRegister, simpleFuncName, serviceType.String(), err)
	}

	return nil
}

// MustResolve resolves a component of the specified generic type T from the container.
//
// It works by:
//  1. Determining the required type (T) via reflection.
//  2. Finding the associated provider using the type as the key.
//  3. Recursively resolving all dependencies required to construct T.
//  4. Performing a type assertion to return the concrete type T, eliminating the need
//     for manual type casting by the caller.
//
// Behavior:
//   - If the component (T) is registered as a Singleton, the same instance is returned on every call.
//   - If the component (T) is registered as Transient, a new instance is created and returned.
//
// Panic Conditions:
//   - Panics if the required component type T is not registered in the container.
//   - Panics if the component or any of its recursive dependencies fails to resolve
//     (e.g., a constructor returns an error).
//
// Parameters:
//   - container: The initialized dependency injection container.
//
// Returns:
// T
func MustResolve[T any](container *Container) T {
	return resolveWithPanic[T](container, "")
}

func MustResolveNamed[T any](container *Container, name string) T {
	return resolveWithPanic[T](container, name)
}

func resolveWithPanic[T any](container *Container, name string) T {
	var zeroT T
	serviceType := reflect.TypeOf(&zeroT).Elem() // Forma más robusta de obtener el tipo T

	providers, found := container.providers[serviceType]
	if !found {
		panic(fmt.Sprintf(errComponentNotRegistered, serviceType.String()))
	}

	if name != "" {
		nameExists := false
		for _, p := range providers {
			if p.name == name {
				nameExists = true
				break
			}
		}
		if !nameExists {
			panic(fmt.Errorf("component %s registered but no implementation found with name: %s", serviceType.String(), name))
		}
	}

	instanceAny, err := container.resolve(serviceType, name)
	if err != nil {
		panic(fmt.Sprintf(errComponentNotResolved, serviceType.String(), err))
	}

	return instanceAny.(T)
}

func (container *Container) registerComponent(constructor any, pType providerType, name string) error {
	originalConstructor := reflect.ValueOf(constructor)
	constructorType := originalConstructor.Type()

	serviceType := constructorType.Out(returnedInstanceIndex)

	dependencies := make([]reflect.Type, constructorType.NumIn())
	for i := 0; i < constructorType.NumIn(); i++ {
		dependencies[i] = constructorType.In(i)
	}

	providerElement := provider{
		constructor:  originalConstructor,
		dependencies: dependencies,
		providerType: pType,
		name:         name,
		instance:     nil,
	}

	if container.providers[serviceType] == nil {
		container.providers[serviceType] = []provider{}
	}

	container.providers[serviceType] = append(container.providers[serviceType], providerElement)

	return nil
}

func (container *Container) resolve(serviceType reflect.Type, name string) (any, error) {
	providers, found := container.providers[serviceType]
	if !found {
		return nil, fmt.Errorf(errNoProvider, serviceType.String())
	}

	var idx int = -1
	for i := range providers {
		if name == "" || providers[i].name == name {
			idx = i
			break
		}
	}

	if idx == -1 {
		return nil, fmt.Errorf("provider not found for name: %s", name)
	}

	p := &container.providers[serviceType][idx]

	switch p.providerType {
	case SINGLETON, APPLICATION:
		if p.instance != nil {
			return p.instance, nil
		}
	case TRANSIENT:
	case REQUEST, SESSION, WEBSOCKET:
		if p.instance != nil {
			return p.instance, nil
		}
	}

	dependencies, err := container.resolveDependencies(p.constructor)
	if err != nil {
		return nil, err
	}

	results := p.constructor.Call(dependencies)

	if len(results) == 2 && !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}

	newInstance := results[0].Interface()

	if p.providerType == SINGLETON || p.providerType == APPLICATION {
		p.instance = newInstance
	}

	return newInstance, nil
}

func (container *Container) resolveDependencies(constructor reflect.Value) ([]reflect.Value, error) {
	t := constructor.Type()
	args := make([]reflect.Value, t.NumIn())

	for i := 0; i < t.NumIn(); i++ {
		depType := t.In(i)

		if depType.Kind() == reflect.Struct {
			newStruct := reflect.New(depType).Elem()
			for j := 0; j < depType.NumField(); j++ {
				field := depType.Field(j)
				tag := field.Tag.Get("inject")

				instance, err := container.resolve(field.Type, tag)
				if err != nil {
					return nil, err
				}
				newStruct.Field(j).Set(reflect.ValueOf(instance))
			}
			args[i] = newStruct
		} else {
			instance, err := container.resolve(depType, "")
			if err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(instance)
		}
	}
	return args, nil
}
