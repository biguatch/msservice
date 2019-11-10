package middleware

import (
	"github.com/biguatch/msjwt"

	"github.com/biguatch/msservice"
)

type Container struct {
	service  msservice.ServiceInterface
	jwtToken msjwt.TokenInterface
}

func NewContainer(service msservice.ServiceInterface, jwtToken msjwt.TokenInterface) *Container {
	container := &Container{
		service:  service,
		jwtToken: jwtToken,
	}

	return container
}
