package container

import "reflect"

type provider struct {
	constructor  reflect.Value
	dependencies []reflect.Type
	name         string
	providerType providerType
	instance     any
}
