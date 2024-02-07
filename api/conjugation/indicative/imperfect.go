package indicative

import (
	"spanish-conjugation-api/api/utils"
)

var imperfectIrregulars = map[string]map[string]string{
	"ser": {
		"yo":       "era",
		"tu":       "eras",
		"usted":    "era",
		"nosotros": "éramos",
		"vosotros": "erais",
		"ustedes":  "eran",
	},
	"ir": {
		"yo":       "iba",
		"tu":       "ibas",
		"usted":    "iba",
		"nosotros": "íbamos",
		"vosotros": "ibais",
		"ustedes":  "iban",
	},
	"ver": {
		"yo":       "veía",
		"tu":       "veías",
		"usted":    "veía",
		"nosotros": "veíamos",
		"vosotros": "veías",
		"ustedes":  "veían",
	},
}

var imperfectEndings = map[string]map[string]string{
	"ar": {
		"yo":       "aba",
		"tu":       "abas",
		"usted":    "aba",
		"nosotros": "ábamos",
		"vosotros": "abais",
		"ustedes":  "aban",
	},
	"er": {
		"yo":       "ía",
		"tu":       "ías",
		"usted":    "ía",
		"nosotros": "íamos",
		"vosotros": "ías",
		"ustedes":  "ían",
	},
	"ir": {
		"yo":       "ía",
		"tu":       "ías",
		"usted":    "ía",
		"nosotros": "íamos",
		"vosotros": "ías",
		"ustedes":  "ían",
	},
}

func GetImperfectEnding(ending string, subject string) string {
	return imperfectEndings[ending][subject]
}

func ConjugateImperfect(verb string, subject string) string {
	irregular, exists := imperfectIrregulars[verb][subject]
	if exists {
		return irregular
	}

	base, ending := utils.SplitVerb(verb)
	newEnding := GetImperfectEnding(ending, subject)

	return base + newEnding
}
