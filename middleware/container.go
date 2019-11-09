package middleware

import "github.com/biguatch/msservice"

type Container struct {
	service *msservice.Service
}

func NewContainer(service *msservice.Service) *Container {
	container := &Container{service: service}

	return container
}
