package router

import (
	"net/http"
	"spanish-conjugation-api/api/conjugation"
	"spanish-conjugation-api/api/middleware"
	"spanish-conjugation-api/api/utils"

	"github.com/go-chi/chi/v5"
)

func Initialize(r *chi.Mux) {

	// serve home page static html
	fs := http.FileServer(http.Dir("./static"))
	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fs.ServeHTTP(res, req)
	})

	r.Route("/{verb}", func(r chi.Router) {
		r.Use(middleware.Text)
		// check if the verb is valid
		// current not going to check if the verb is actually a spanish verb, rather just ensure that it ends in ar/er/ir
		r.Use(middleware.Verb)

		// handle all verbs besides imperative mood
		r.Get("/{tense}/{mood}/{subject}", func(res http.ResponseWriter, req *http.Request) {
			verb := chi.URLParam(req, "verb")
			tense := chi.URLParam(req, "tense")
			mood := chi.URLParam(req, "mood")
			subject := chi.URLParam(req, "subject")

			// error handling
			allowedMoods := []string{"indicative", "subjunctive"}
			if !utils.ArrayContainsString(allowedMoods, mood) {
				http.Error(res, "The given mood is not valid, use indicative, subjunctive, or imperative.", http.StatusBadRequest)
				return
			}

			allowedSubjects := []string{"yo", "nosotros", "tu", "vosotros", "usted", "ustedes"}
			if !utils.ArrayContainsString(allowedSubjects, subject) {
				http.Error(res, "The given subject is not valid, use yo, nosotros, tu, vosotros, usted, or ustedes.", http.StatusBadRequest)
				return
			}

			if mood == "indicative" {
				allowedTensesForIndicative := []string{"present", "nearfuture", "preterite", "imperfect", "future", "conditional", "presentperfect", "futureperfect", "pluperfect", "conditionalperfect", "conditionalprogressive", "futureprogressive", "presentprogressive", "pastprogressive"}
				if !utils.ArrayContainsString(allowedTensesForIndicative, tense) {
					http.Error(res, "The given tense is not valid for the indicative mood.", http.StatusBadRequest)
					return
				}
			}
			if mood == "subjunctive" {
				allowedTensesForSubjunctive := []string{"present", "imperfect", "presentperfect", "pluperfect"}
				if !utils.ArrayContainsString(allowedTensesForSubjunctive, tense) {
					http.Error(res, "The given tense is not valid for the subjunctive mood.", http.StatusBadRequest)
					return
				}
			}

			// conjugation
			conjugated := conjugation.Conjugate(verb, mood, tense, subject)

			res.WriteHeader(http.StatusOK)
			res.Write([]byte(conjugated))
		})

		// handle imperative mood, which does not have a tense
		r.Get("/{mood}/{subject}", func(res http.ResponseWriter, req *http.Request) {
			verb := chi.URLParam(req, "verb")
			mood := chi.URLParam(req, "mood")
			subject := chi.URLParam(req, "subject")

			allowedMoods := []string{"imperative"}
			if !utils.ArrayContainsString(allowedMoods, mood) {
				http.Error(res, "The given mood is not valid, use indicative, subjunctive, or imperative.", http.StatusBadRequest)
				return
			}

			allowedSubjects := []string{"nosotros", "tu", "vosotros", "usted", "ustedes"}
			if !utils.ArrayContainsString(allowedSubjects, subject) {
				http.Error(res, "The given subject is not valid, for imperative mood use nosotros, tu, vosotros, usted, or ustedes.", http.StatusBadRequest)
				return
			}

			// conjugation
			conjugated := conjugation.Conjugate(verb, mood, "", subject)

			res.WriteHeader(http.StatusOK)
			res.Write([]byte(conjugated))
		})
	})
}
