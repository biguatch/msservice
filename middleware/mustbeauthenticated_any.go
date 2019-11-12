package middleware

import (
	"errors"
	"net/http"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeAuthenticatedAny(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		service2service := r.Context().Value("service2service")
		if user == nil && service2service != true {
			msservice.SendError(w, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
