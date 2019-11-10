package middleware

import (
	"net"
	"net/http"
	"time"
)

func (container *Container) RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		container.service.GetLogger().Logrus().
			WithField("RequestTime", time.Now()).
			WithField("RequestMethod", r.Method).
			WithField("RequestUrl", r.URL.String()).
			WithField("UserAgent", r.UserAgent()).
			WithField("Referrer", r.Referer()).
			WithField("Proto", r.Proto).
			WithField("RemoteIP", ipFromHostPort(r.RemoteAddr)).
			Info()
		next.ServeHTTP(w, r)
	})
}

func ipFromHostPort(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h
}
