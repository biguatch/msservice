package middleware

import (
	"errors"
	"net/http"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeAdminOrService(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service2service := r.Context().Value("service2service")
		user := r.Context().Value("user")
		isAdmin := false

		if user != nil {
			user := user.(msservice.Identity)
			isAdmin = user.IsAdmin
		}

		if (service2service == nil || !service2service.(bool)) && !isAdmin {
			msservice.SendError(w, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
