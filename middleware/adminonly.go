package middleware

import (
	"net/http"

	"github.com/biguatch/msservice"
)

func (container *Container) AdminOnly(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") == nil {
			resp := &msservice.Response{
				Success: false,
				Error: &msservice.Error{
					Message: "unauthorized",
				},
			}
			container.service.Respond(w, *resp, http.StatusUnauthorized)
			return
		}

		user := r.Context().Value("user").(msservice.Identity)

		if !user.IsAdmin {
			resp := &msservice.Response{
				Success: false,
				Error: &msservice.Error{
					Message: "unauthorized",
				},
			}
			container.service.Respond(w, *resp, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
