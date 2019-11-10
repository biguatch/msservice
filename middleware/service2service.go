package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
)

func CreateService2Service(localSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Service-2-Service") == "" {
				next.ServeHTTP(w, r)
				return
			}
			incoming := r.Header.Get("Service-2-Service")

			hm := hmac.New(sha256.New, []byte(localSecret))
			hm.Write([]byte(r.URL.Path))
			calculated := base64.StdEncoding.EncodeToString(hm.Sum(nil))

			if incoming != calculated {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "service2service", true)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
