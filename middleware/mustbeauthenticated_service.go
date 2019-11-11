package middleware

import (
	"net/http"

	"github.com/biguatch/msutil"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeServiceAuthenticated(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service2service := r.Context().Value("service2service")
		if msutil.CheckBoolen(service2service) != true {
			resp := &msservice.Response{
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
