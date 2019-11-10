package middleware

import (
	"net/http"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeAuthenticatedAny(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		service2service := r.Context().Value("service2service")
		if user == nil && service2service != true {
			resp := &msservice.Response{
				Success: false,
				Error: &msservice.Error{
					Message: "unauthorized",
				},
			}
			resp.SendResponse(w, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
