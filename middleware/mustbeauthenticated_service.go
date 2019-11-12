package middleware

import (
	"errors"
	"net/http"

	"github.com/biguatch/msutil"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeServiceAuthenticated(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service2service := r.Context().Value("service2service")
		if msutil.CheckBoolen(service2service) != true {
			msservice.SendError(w, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
