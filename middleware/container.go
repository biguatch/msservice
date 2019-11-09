package middleware

import (
	"github.com/biguatch/msjwt"

	"github.com/biguatch/msservice"
)

type Container struct {
	service  *msservice.Service
	jwtToken *msjwt.Token
}

func NewContainer(service *msservice.Service, jwtToken *msjwt.Token) *Container {
	container := &Container{
		service:  service,
		jwtToken: jwtToken,
	}

	return container
}
