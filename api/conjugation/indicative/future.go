package indicative

import (
	"spanish-conjugation-api/api/utils"
)

func GetFutureStem(verb string) string {
	// decir
	if verb == "decir" {
		return "dir"
	}

	// hacer
	if verb == "hacer" {
		return "har"
	}

	// vowel -> d
	if utils.HasOneOfMultipleSuffixes(verb, "tener", "poner", "valer", "salir", "venir") {
		return verb[:len(verb)-2] + "dr"
	}

	// remove vowel
	if utils.HasOneOfMultipleSuffixes(verb, "poder", "caber", "querer", "saber", "haber") {
		return verb[:len(verb)-2] + "r"
	}

	return verb
}

var futureEndings = map[string]string{
	"yo":       "é",
	"tu":       "ás",
	"usted":    "á",
	"nosotros": "emos",
	"vosotros": "éis",
	"ustedes":  "án",
}

func ConjugateFuture(verb string, subject string) string {
	stem := GetFutureStem(verb)
	ending, exists := futureEndings[subject]
	if !exists {
		return "Invalid subject given."
	}
	return stem + ending
}
