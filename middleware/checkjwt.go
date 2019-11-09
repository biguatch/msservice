package middleware

import (
	"context"
	"net/http"
	"strings"
)

func (container *Container) CheckJwt(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.NotFound(w, r)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.NotFound(w, r)
			return
		}

		claims, err := container.jwtToken.Validate(parts[1])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", struct {
			Id      string `json:"id"`
			IsAdmin bool   `json:"is_admin"`
		}{
			Id:      claims["sub"].(string),
			IsAdmin: claims["is_admin"].(bool),
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
