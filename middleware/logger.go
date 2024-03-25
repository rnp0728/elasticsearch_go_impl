package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// Logger is a middleware function that logs the request and response information.
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // start time is the time when the request is received by the server
        start := time.Now()

        // create a new wrapped writer that wraps the original response writer
        wrapped := &wrappedWriter{
            ResponseWriter: w,
            statusCode:     http.StatusOK,
        }

        // call the next handler in the middleware chain
        next.ServeHTTP(wrapped, r)

        // log the request and response information
        log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
    })
}
