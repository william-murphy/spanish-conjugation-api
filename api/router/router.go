package router

import (
	"net/http"
	"spanish-conjugation-api/api/conjugation"
	"spanish-conjugation-api/api/middleware"

	"github.com/go-chi/chi/v5"
)

func arrayContainsString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

func Initialize(r *chi.Mux) {
	r.Use(middleware.Text)
	// check if the verb is valid
	// current not going to check if the verb is actually a spanish verb, rather just ensure that it ends in ar/er/ir
	r.Use(middleware.VerifyVerb)

	// handle all verbs besides imperative mood
	r.Get("/{verb}/{tense}/{mood}/{subject}", func(res http.ResponseWriter, req *http.Request) {
		verb := chi.URLParam(req, "verb")
		tense := chi.URLParam(req, "tense")
		mood := chi.URLParam(req, "mood")
		subject := chi.URLParam(req, "subject")

		// error handling
		allowedMoods := []string{"indicative", "subjunctive"}
		if !arrayContainsString(allowedMoods, mood) {
			http.Error(res, "The given mood is not valid, use indicative, subjunctive, or imperative.", http.StatusBadRequest)
			return
		}

		allowedSubjects := []string{"yo", "nosotros", "tu", "vosotros", "usted", "ustedes"}
		if !arrayContainsString(allowedSubjects, subject) {
			http.Error(res, "The given subject is not valid, use yo, nosotros, tu, vosotros, usted, or ustedes.", http.StatusBadRequest)
			return
		}

		allowedTensesForIndicative := []string{"present", "nearfuture", "preterite", "imperfect", "future", "conditional", "presentperfect", "futureperfect", "pluperfect", "conditionalperfect", "conditionalprogressive", "futureprogressive", "presentprogressive", "pastprogressive"}
		if mood == "indicative" {
			if !arrayContainsString(allowedTensesForIndicative, tense) {
				http.Error(res, "The given tense is not valid for the indicative mood.", http.StatusBadRequest)
				return
			}
		}
		allowedTensesForSubjunctive := []string{"present", "imperfect", "presentperfect", "pluperfect"}
		if mood == "subjunctive" {
			if !arrayContainsString(allowedTensesForSubjunctive, tense) {
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
	r.Get("/{verb}/{mood}/{subject}", func(res http.ResponseWriter, req *http.Request) {
		verb := chi.URLParam(req, "verb")
		mood := chi.URLParam(req, "mood")
		subject := chi.URLParam(req, "subject")

		allowedMoods := []string{"imperative"}
		if !arrayContainsString(allowedMoods, mood) {
			http.Error(res, "The given mood is not valid, use indicative, subjunctive, or imperative.", http.StatusBadRequest)
			return
		}

		allowedSubjects := []string{"nosotros", "tu", "vosotros", "usted", "ustedes"}
		if !arrayContainsString(allowedSubjects, subject) {
			http.Error(res, "The given subject is not valid, for imperative mood use nosotros, tu, vosotros, usted, or ustedes.", http.StatusBadRequest)
			return
		}

		// conjugation
		conjugated := conjugation.Conjugate(verb, mood, "", subject)

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(conjugated))
	})
}
