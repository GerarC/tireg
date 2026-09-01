package container

import "reflect"

type providerMap map[reflect.Type][]provider

type providerType int

const (
	SINGLETON   providerType = iota
	TRANSIENT   providerType = iota
	REQUEST     providerType = iota
	SESSION     providerType = iota
	APPLICATION providerType = iota
	WEBSOCKET   providerType = iota
)
