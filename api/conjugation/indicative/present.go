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
	"ver":   "veo",
	"caber": "quepo",
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

func getIrregularPresentYoStem(verb string, base string) string {
	// quir verbs
	if strings.HasSuffix(verb, "quir") {
		return utils.ChangeStem(base, "qu", "c")
	}

	// -ger and -gir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "ger", "gir") {
		return base[:len(base)-1] + "j"
	}

	// -eguir verbs
	if strings.HasSuffix(verb, "eguir") {
		return base[:len(base)-3] + "ig"
	}

	// -guir verbs
	if strings.HasSuffix(verb, "guir") {
		return base[:len(base)-1]
	}

	// -igo verbs
	if strings.HasSuffix(verb, "decir") {
		return base[:len(base)-2] + "ig"
	}

	// -ago verbs
	if utils.HasOneOfMultipleSuffixes(verb, "hacer", "facer") {
		return base[:len(base)-1] + "g"
	}

	// -asgo verbs
	if strings.HasSuffix(verb, "asir") {
		return base + "g"
	}

	// -algo verbs
	if utils.HasOneOfMultipleSuffixes(verb, "salir", "valer") {
		return base + "g"
	}

	// -oir and -aer verbs
	if utils.HasOneOfMultipleSuffixes(verb, "oir", "aer") {
		return base + "ig"
	}

	// -engo and -ongo verbs
	if utils.HasOneOfMultipleSuffixes(verb, "tener", "poner") {
		return base + "g"
	}

	// consonant + -cer and -cir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "cer", "cir") && !utils.ArrayContainsString(utils.Vowels, base[:len(base)-2]) {
		return base[:len(base)-1] + "z"
	}

	// vowel + -cer and -cir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "cer", "cir") && utils.ArrayContainsString(utils.Vowels, base[:len(base)-2]) {
		return base[:len(base)-1] + "zc"
	}

	return getIrregularPresentStem(verb, base)
}

func getIrregularPresentStem(verb string, base string) string {
	// -uir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "oir", "uir") {
		return base + "y"
	}

	// oler
	if verb == "oler" {
		return utils.ChangeStem(base, "o", "hue")
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
		return getIrregularPresentYoStem(verb, base)
	} else if subject == "nosotros" || subject == "vosotros" {
		return base
	} else {
		return getIrregularPresentStem(verb, base)
	}
}

func GetPresentEnding(ending string, subject string) string {
	return presentEndings[ending][subject]
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
	newEnding := GetPresentEnding(ending, subject)

	return stem + newEnding
}
