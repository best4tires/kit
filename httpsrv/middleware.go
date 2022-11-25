package httpsrv

import (
	"net/http"
	"time"

	"github.com/best4tires/kit/log"
	"github.com/gorilla/handlers"
)

func GZIP() func(http.Handler) http.Handler {
	return handlers.CompressHandler
}

func Logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sw := NewStatusWriter(w)
			t0 := time.Now()
			next.ServeHTTP(sw, r)
			log.Accessf("%s host=%q path=%q query=%q => status %d (%s) in %s",
				r.Method, r.Host, r.URL.Path, r.URL.RawQuery, sw.statusCode, http.StatusText(sw.statusCode), time.Since(t0))
		})
	}
}

func Cors() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-control-allow-origin", "*")
			w.Header().Add("Access-control-allow-methods", "*")
			w.Header().Add("Access-control-allow-headers", "*")
			next.ServeHTTP(w, r)
		})
	}
}

func Recovery() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Errorf("http-request: recovered: %v", err)
					log.DebugStack()
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
