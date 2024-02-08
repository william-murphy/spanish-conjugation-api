package imperative

import (
	"spanish-conjugation-api/api/conjugation/indicative"
	"spanish-conjugation-api/api/conjugation/subjunctive"
	"spanish-conjugation-api/api/utils"
	"strings"
)

func ConjugateTu(verb string) string {
	// irregulars
	if utils.HasOneOfMultipleSuffixes(verb, "tener", "salir", "poner", "venir") {
		return verb[:len(verb)-2]
	}
	if strings.HasSuffix(verb, "hacer") {
		return "haz"
	}
	if verb == "decir" {
		return "di"
	}
	if verb == "ser" {
		return "sé"
	}
	if verb == "ir" {
		return "ve"
	}
	return indicative.ConjugatePresent(verb, "usted")
}

func ConjugateVosotros(verb string) string {
	if verb == "ir" {
		// conflicting opinions on if this is irregular or not... gonna stick with the irregular
		return "idos"
	}
	return verb[:len(verb)-1] + "d"
}

func ConjugateUsted(verb string) string {
	return subjunctive.ConjugatePresent(verb, "usted")
}

func ConjugateUstedes(verb string) string {
	return subjunctive.ConjugatePresent(verb, "ustedes")
}

func ConjugateNosotros(verb string) string {
	return subjunctive.ConjugatePresent(verb, "nosotros")
}
