package middleware

import (
	"net/http"
)

// Auth - Method Definition
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
