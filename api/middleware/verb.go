package middleware

import (
	"net/http"
	"spanish-conjugation-api/api/utils"

	"github.com/go-chi/chi/v5"
)

func Verb(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		verb := chi.URLParam(req, "verb")

		if !utils.ContainsOnlyLowerCase(verb) {
			http.Error(res, "Spanish verbs must contain only lowercase letters a-z.", http.StatusBadRequest)
			return
		}

		if !utils.VerifyVerbEnding(verb) {
			http.Error(res, "Spanish verbs must end in ar, er, or ir.", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(res, req)
	})
}
