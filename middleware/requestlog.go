package middleware

import (
	"net"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func (container *Container) RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqTime := time.Now()
		sw := statusWriter{ResponseWriter: w}

		next.ServeHTTP(&sw, r)

		resTime := time.Now()
		container.service.GetLogger().Logrus().
			WithField("RequestTime", reqTime).
			WithField("ResponseTime", resTime).
			WithField("Duration", resTime.Sub(reqTime)).
			WithField("ResponseStatus", sw.status).
			WithField("ContentLength", sw.length).
			WithField("RequestMethod", r.Method).
			WithField("RequestUrl", r.URL.String()).
			WithField("UserAgent", r.UserAgent()).
			WithField("Referrer", r.Referer()).
			WithField("Proto", r.Proto).
			WithField("RemoteIP", ipFromHostPort(r.RemoteAddr)).
			Info()
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
