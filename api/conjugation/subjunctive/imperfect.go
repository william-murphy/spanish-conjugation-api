package subjunctive

import (
	"spanish-conjugation-api/api/conjugation/indicative"
)

var imperfectEndings = map[string]string{
	"yo":       "ra",
	"tu":       "ras",
	"usted":    "ra",
	"nosotros": "ramos",
	"vosotros": "rais",
	"ustedes":  "ran",
}

func ConjugateImperfect(verb string, subject string) string {
	preteriteConjugation := indicative.ConjugatePreterite(verb, "ustedes")
	stem := preteriteConjugation[:len(preteriteConjugation)-3]

	// last vowel before ending needs accent mark if subject is nosotros
	if subject == "nosotros" {
		if stem[len(stem)-1] == 'a' {
			stem = stem[:len(stem)-1] + "á"
		} else if stem[len(stem)-1] == 'e' {
			stem = stem[:len(stem)-1] + "é"
		}
	}

	return stem + imperfectEndings[subject]
}
