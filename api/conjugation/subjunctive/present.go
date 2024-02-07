package subjunctive

import (
	"spanish-conjugation-api/api/conjugation/indicative"
	"spanish-conjugation-api/api/utils"
	"strings"
)

var presentIrregulars = map[string]map[string]string{
	"ser": {
		"yo":       "sea",
		"tu":       "seas",
		"usted":    "sea",
		"nosotros": "seamos",
		"vosotros": "seáis",
		"ustedes":  "sean",
	},
	"ir": {
		"yo":       "vaya",
		"tu":       "vayas",
		"usted":    "vaya",
		"nosotros": "vayamos",
		"vosotros": "vayáis",
		"ustedes":  "vayan",
	},
	"estar": {
		"yo":       "esté",
		"tu":       "estés",
		"usted":    "esté",
		"nosotros": "estemos",
		"vosotros": "estéis",
		"ustedes":  "estén",
	},
	"saber": {
		"yo":       "sepa",
		"tu":       "sepas",
		"usted":    "sepa",
		"nosotros": "sepamos",
		"vosotros": "sepáis",
		"ustedes":  "sepan",
	},
	"haber": {
		"yo":       "haya",
		"tu":       "hayas",
		"usted":    "haya",
		"nosotros": "hayamos",
		"vosotros": "hayáis",
		"ustedes":  "hayan",
	},
	"dar": {
		"yo":       "dé",
		"tu":       "des",
		"usted":    "dé",
		"nosotros": "demos",
		"vosotros": "déis",
		"ustedes":  "den",
	},
}

var presentEndings = map[string]map[string]string{
	"ar": {
		"yo":       "e",
		"tu":       "es",
		"usted":    "e",
		"nosotros": "emos",
		"vosotros": "éis",
		"ustedes":  "en",
	},
	"er": {
		"yo":       "a",
		"tu":       "as",
		"usted":    "a",
		"nosotros": "amos",
		"vosotros": "áis",
		"ustedes":  "an",
	},
	"ir": {
		"yo":       "a",
		"tu":       "as",
		"usted":    "a",
		"nosotros": "amos",
		"vosotros": "áis",
		"ustedes":  "an",
	},
}

func getPresentStem(verb string, base string, subject string) string {
	// TODO - handle the case of ir verb stem changes persisting in nostros/vosotros
	// link: https://conjuguemos.com/tenses/spanish/12/47
	stem, irregular := indicative.GetIrregularPresentYoStem(verb, base)
	if !irregular && subject != "nosotros" && subject != "vosotros" {
		stem = indicative.GetIrregularPresentStem(verb, base)
	}

	if strings.HasSuffix(verb, "car") {
		stem = utils.ChangeStem(stem, "c", "qu")
	} else if strings.HasSuffix(verb, "gar") {
		stem = utils.ChangeStem(stem, "g", "gu")
	} else if strings.HasSuffix(verb, "zar") {
		stem = utils.ChangeStem(stem, "z", "c")
	}

	return stem
}

func ConjugatePresent(verb string, subject string) string {
	irregular, exists := presentIrregulars[verb][subject]
	if exists {
		return irregular
	}

	base, ending := utils.SplitVerb(verb)

	newEnding := presentEndings[ending][subject]
	stem := getPresentStem(verb, base, subject)

	return stem + newEnding
}
