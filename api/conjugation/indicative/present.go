package indicative

import (
	"spanish-conjugation-api/api/utils"
	"strings"
)

var presentIrregulars = map[string]map[string]string{
	"ser": {
		"yo":       "soy",
		"tu":       "eres",
		"usted":    "es",
		"nosotros": "somos",
		"vosotros": "soi",
		"ustedes":  "son",
	},
	"ir": {
		"yo":       "voy",
		"tu":       "vas",
		"usted":    "va",
		"nosotros": "vamos",
		"vosotros": "vais",
		"ustedes":  "van",
	},
	"haber": {
		"yo":       "he",
		"tu":       "has",
		"usted":    "ha",
		"nosotros": "hemos",
		"vosotros": "habeis",
		"ustedes":  "han",
	},
	"estar": {
		"yo":       "estoy",
		"tu":       "estás",
		"usted":    "está",
		"nosotros": "estamos",
		"vosotros": "estáis",
		"ustedes":  "están",
	},
}

var irregularYo = map[string]string{
	"saber": "se",
	"dar":   "doy",
}

var presentEndings = map[string]map[string]string{
	"ar": {
		"yo":       "o",
		"tu":       "as",
		"usted":    "a",
		"nosotros": "amos",
		"vosotros": "áis",
		"ustedes":  "an",
	},
	"er": {
		"yo":       "o",
		"tu":       "es",
		"usted":    "e",
		"nosotros": "emos",
		"vosotros": "éis",
		"ustedes":  "en",
	},
	"ir": {
		"yo":       "o",
		"tu":       "es",
		"usted":    "e",
		"nosotros": "imos",
		"vosotros": "ís",
		"ustedes":  "en",
	},
}

func GetIrregularPresentYoStem(verb string, base string) (string, bool) {
	// ver
	if verb == "ver" {
		return "ve", true
	}

	// caber
	if verb == "caber" {
		return "quep", true
	}

	// reir
	if strings.HasSuffix(verb, "reir") {
		return utils.ChangeStem(base, "e", "í"), true
	}

	// quir verbs
	if strings.HasSuffix(verb, "quir") {
		return utils.ChangeStem(base, "qu", "c"), true
	}

	// -ger and -gir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "ger", "gir") {
		return base[:len(base)-1] + "j", true
	}

	// -eguir verbs
	if strings.HasSuffix(verb, "eguir") {
		return base[:len(base)-3] + "ig", true
	}

	// -guir verbs
	if strings.HasSuffix(verb, "guir") {
		return base[:len(base)-1], true
	}

	// -uir verbs
	if strings.HasSuffix(verb, "uir") {
		return base + "y", true
	}

	// -igo verbs
	if strings.HasSuffix(verb, "decir") {
		return base[:len(base)-2] + "ig", true
	}

	// -ago verbs
	if utils.HasOneOfMultipleSuffixes(verb, "hacer", "facer") {
		return base[:len(base)-1] + "g", true
	}

	// -go verbs
	if utils.HasOneOfMultipleSuffixes(verb, "salir", "valer", "tener", "poner", "asir") {
		return base + "g", true
	}

	// -oir and -aer verbs
	if utils.HasOneOfMultipleSuffixes(verb, "oir", "aer") {
		return base + "ig", true
	}

	// consonant + -cer and -cir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "cer", "cir") && !utils.ArrayContainsString(utils.Vowels, string(base[len(base)-2])) {
		return base[:len(base)-1] + "z", true
	}

	// vowel + -cer and -cir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "cer", "cir") && utils.ArrayContainsString(utils.Vowels, string(base[len(base)-2])) {
		return base[:len(base)-1] + "zc", true
	}

	return base, false
}

func GetIrregularPresentStem(verb string, base string) string {
	// oler
	if verb == "oler" {
		return utils.ChangeStem(base, "o", "hue")
	}

	// -uir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "oir", "uir") {
		return base + "y"
	}

	// e -> ie
	if utils.HasOneOfMultipleSuffixes(verb, "certar", "vesar", "cerrar", "fesar", "pertar", "pezar", "regar", "nevar", "mendar", "sentar", "lentar", "menzar", "helar", "bernar", "negar", "pensar", "querer", "fender", "perder", "tender", "cender", "tener", "mentir", "gerir", "vertir", "ferir", "sentir", "venir") {
		return utils.ChangeStem(base, "e", "ie")
	}

	// e -> i
	if utils.HasOneOfMultipleSuffixes(verb, "eguir", "decir", "elegir", "medir", "reir", "servir", "vestir", "regir", "pedir", "petir") {
		return utils.ChangeStem(base, "e", "i")
	}

	// o -> ue
	if utils.HasOneOfMultipleSuffixes(verb, "sonar", "morzar", "costar", "mostrar", "probar", "cordar", "ronar", "volar", "rogar", "contrar", "contar", "colgar", "solver", "cocer", "mover", "volver", "morder", "soler", "torcer", "poder", "moler", "llover", "doler", "dormir", "morir") {
		return utils.ChangeStem(base, "o", "ue")
	}

	// i -> ie
	if strings.HasSuffix(verb, "quirir") {
		return utils.ChangeStem(base, "i", "ie")
	}

	// u -> ue
	if strings.HasSuffix(verb, "jugar") {
		return utils.ChangeStem(base, "u", "ue")
	}

	return base
}

func GetPresentStem(verb string, base string, subject string) string {
	if subject == "yo" {
		stem, irregular := GetIrregularPresentYoStem(verb, base)
		if irregular {
			return stem
		}
	} else if subject == "nosotros" || subject == "vosotros" {
		return base
	}
	return GetIrregularPresentStem(verb, base)
}

func ConjugatePresent(verb string, subject string) string {
	irregular, exists := presentIrregulars[verb][subject]
	if exists {
		return irregular
	}

	value, exists := irregularYo[verb]
	if subject == "yo" && exists {
		return value
	}

	base, ending := utils.SplitVerb(verb)

	stem := GetPresentStem(verb, base, subject)
	newEnding := presentEndings[ending][subject]

	return stem + newEnding
}
