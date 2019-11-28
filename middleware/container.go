package middleware

import (
	"net/http"

	"github.com/biguatch/msjwt"

	"github.com/biguatch/msservice"
)

type Container struct {
	service           msservice.ServiceInterface
	jwtToken          msjwt.TokenInterface
	requestLogFilters []func(r *http.Request) bool
}

func NewContainer(service msservice.ServiceInterface, jwtToken msjwt.TokenInterface) *Container {
	container := &Container{
		service:  service,
		jwtToken: jwtToken,
	}

	return container
}

func (container *Container) AddRequestLogFilter(f func(r *http.Request) bool) {
	container.requestLogFilters = append(container.requestLogFilters, f)
}
