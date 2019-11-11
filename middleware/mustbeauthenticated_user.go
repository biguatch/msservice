package middleware

import (
	"net/http"

	"github.com/biguatch/msservice"
)

func (container *Container) MustBeUserAuthenticated(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") == nil {
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
