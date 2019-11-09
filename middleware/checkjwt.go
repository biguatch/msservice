package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/biguatch/msservice"
)

func (container *Container) CheckJwt(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			next.ServeHTTP(w, r)
			return
		}

		claims, err := container.jwtToken.Validate(parts[1])
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", msservice.Identity{
			Id:      claims["sub"].(string),
			IsAdmin: claims["is_admin"].(bool),
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
