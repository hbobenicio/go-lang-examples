package middleware

import (
	"fmt"
	"net/http"
)

// LogRequest logs requests (duh!)
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method, req.URL.Path)
		next.ServeHTTP(res, req)
	})
}
