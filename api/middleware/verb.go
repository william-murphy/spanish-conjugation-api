package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func VerifyVerb(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		verb := chi.URLParam(req, "verb")

		if !strings.HasSuffix(verb, "ar") || !strings.HasSuffix(verb, "er") || !strings.HasSuffix(verb, "ir") {
			http.Error(res, "Spanish verbs must end in ar, er, or ir.", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(res, req)
	})
}
