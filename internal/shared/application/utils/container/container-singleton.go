package container

import (
	"sync"
)

var (
	instance *Container
	once     sync.Once
)

func initializeContainer() {
	once.Do(func() {
		instance = NewContainer()
	})
}

func GetInstance() *Container {
	if instance == nil {
		initializeContainer()
	}
	return instance
}
