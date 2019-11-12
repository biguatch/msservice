package middleware

import (
	"errors"
	"net/http"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeUserAuthenticated(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") == nil {
			msservice.SendError(w, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
