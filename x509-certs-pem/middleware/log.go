package middleware

import (
	"log"
	"net/http"
	"time"
)

// LogRequest is a simple middleware that logs the spent time for handling the request
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)
		log.Println(r.Method, r.URL, elapsed.Truncate(time.Millisecond))
	})
}
