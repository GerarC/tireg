package container

import (
	"fmt"
	"reflect"
)

func validateConstructorType(container *Container, constructorType reflect.Type) error {
	if constructorType.Kind() != reflect.Func {
		return fmt.Errorf(errMustBeFunction, constructorType.Kind())
	}
	if constructorType.NumOut() == emptyReturnedList || constructorType.NumOut() > constructorReturnedElementsLimit {
		return fmt.Errorf(errInvalidOutputs, constructorType.String())
	}
	if constructorType.NumOut() == constructorReturnedElementsLimit && constructorType.Out(returnedErrorIndex).String() != errorTypeName {
		return fmt.Errorf(errSecondMustBeError, constructorType.String())
	}

	serviceType := constructorType.Out(returnedInstanceIndex)
	if _, exists := container.providers[serviceType]; exists {
		return fmt.Errorf(errConstructorTypeDup, serviceType.String())
	}

	return nil
}
