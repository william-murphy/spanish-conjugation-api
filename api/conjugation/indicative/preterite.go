package indicative

import (
	"spanish-conjugation-api/api/utils"
	"strings"
)

var irregularPreteriteEndings = map[string]string{
	"yo":       "e",
	"tu":       "isted",
	"usted":    "o",
	"nosotros": "imos",
	"vosotros": "isteis",
	"ustedes":  "ieron",
}

var preteriteEndings = map[string]map[string]string{
	"ar": {
		"yo":       "é",
		"tu":       "aste",
		"usted":    "ó",
		"nosotros": "amos",
		"vosotros": "asteis",
		"ustedes":  "aron",
	},
	"er": {
		"yo":       "í",
		"tu":       "iste",
		"usted":    "ió",
		"nosotros": "imos",
		"vosotros": "isteis",
		"ustedes":  "ieron",
	},
	"ir": {
		"yo":       "í",
		"tu":       "iste",
		"usted":    "ió",
		"nosotros": "imos",
		"vosotros": "isteis",
		"ustedes":  "ieron",
	},
}

var preteriteFullyIrregulars = map[string]map[string]string{
	"ser": {
		"yo":       "fui",
		"tu":       "fuiste",
		"usted":    "fue",
		"nosotros": "fuimos",
		"vosotros": "fuisteis",
		"ustedes":  "fueron",
	},
	"ir": {
		"yo":       "fui",
		"tu":       "fuiste",
		"usted":    "fue",
		"nosotros": "fuimos",
		"vosotros": "fuisteis",
		"ustedes":  "fueron",
	},
	"dar": {
		"yo":       "di",
		"tu":       "diste",
		"usted":    "dio",
		"nosotros": "dimos",
		"vosotros": "disteis",
		"ustedes":  "dieron",
	},
}

var preteriteIrregularSuffixes = map[string]string{
	"venir": "vin",
	"poner": "pus",
	"tener": "tuv",
	"hacer": "hic",
	"facer": "fic",
	"decir": "dij",
	"traer": "traj",
	"ucir":  "uj",
}

var preteriteIrregulars = map[string]string{
	"andar":  "anduv",
	"estar":  "estuv",
	"poder":  "pud",
	"saber":  "sup",
	"querer": "quis",
	"caber":  "cup",
	"haber":  "hub",
}

func getIrregularPreteriteYoStem(verb string, base string) string {
	// car -> qu
	if strings.HasSuffix(verb, "car") {
		return base[:len(base)-1] + "qu"
	}

	// gar -> gu
	if strings.HasSuffix(verb, "gar") {
		return base[:len(base)-1] + "gu"
	}

	// zar -> c
	if strings.HasSuffix(verb, "zar") {
		return base[:len(base)-1] + "c"
	}

	return base
}

func getIrregularPreteriteBasementStem(verb string, base string) string {
	// e -> i
	if utils.HasOneOfMultipleSuffixes(verb, "eguir", "elegir", "medir", "reir", "servir", "vestir", "regir", "pedir", "petir", "mentir", "gerir", "vertir", "ferir", "sentir") {
		return utils.ChangeStem(base, "e", "i")
	}

	// o -> u
	if utils.HasOneOfMultipleSuffixes(verb, "dormir", "morir") {
		return utils.ChangeStem(base, "o", "u")
	}

	// i -> y
	if utils.HasOneOfMultipleSuffixes(verb, "aer", "eer", "oir", "oer") {
		return base + "y"
	}

	return base
}

func getPreteriteStem(verb string, base string, subject string) (string, bool) {
	// check fully irregulars
	stem, exists := preteriteIrregulars[verb]
	if exists {
		return stem, true
	}

	// check irregular suffixes
	for key, value := range preteriteIrregularSuffixes {
		if strings.HasSuffix(verb, key) {
			newStem := verb[:len(verb)-len(key)] + value
			// special case for hacer and facer verbs
			if (key == "hacer" || key == "facer") && subject == "usted" {
				newStem = newStem[:len(newStem)-1] + "z"
			}
			return newStem, true
		}
	}

	// check irregular in yo form (if subject == "yo")
	if subject == "yo" {
		return getIrregularPreteriteYoStem(verb, base), false
	}

	// check irregular in basement form (if subject == "usted/ustedes")
	if subject == "usted" || subject == "ustedes" {
		return getIrregularPreteriteBasementStem(verb, base), false
	}

	return base, false
}

func getPreteriteEnding(stem string, ending string, subject string, irregular bool) string {
	if irregular {
		if subject == "ustedes" && strings.HasSuffix(stem, "j") {
			return "eron"
		}
		return irregularPreteriteEndings[subject]
	}

	if (subject == "usted" || subject == "ustedes") && strings.HasSuffix(stem, "y") {
		return preteriteEndings[ending][subject][1:]
	}

	return preteriteEndings[ending][subject]
}

func ConjugatePreterite(verb string, subject string) string {
	value, exists := preteriteFullyIrregulars[verb][subject]
	if exists {
		return value
	}

	base, ending := utils.SplitVerb(verb)

	stem, irregular := getPreteriteStem(verb, base, subject)
	newEnding := getPreteriteEnding(stem, ending, subject, irregular)

	return stem + newEnding
}
