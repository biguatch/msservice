package middleware

import (
	"net/http"
)

func (container *Container) Authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") == nil {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}
}
