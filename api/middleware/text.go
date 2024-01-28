package middleware

import (
	"net/http"
)

func Text(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		next.ServeHTTP(res, req)
	})
}
