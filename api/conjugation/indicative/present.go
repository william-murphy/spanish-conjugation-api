package indicative

import (
	"spanish-conjugation-api/api/utils"
	"strings"
)

var ser = map[string]string{
	"yo":       "soy",
	"tu":       "eres",
	"usted":    "es",
	"nosotros": "somos",
	"vosotros": "soi",
	"ustedes":  "son",
}

var haber = map[string]string{
	"yo":       "he",
	"tu":       "has",
	"usted":    "ha",
	"nosotros": "hemos",
	"vosotros": "habeis",
	"ustedes":  "han",
}

var ir = map[string]string{
	"yo":       "voy",
	"tu":       "vas",
	"usted":    "va",
	"nosotros": "vamos",
	"vosotros": "vais",
	"ustedes":  "van",
}

var irregularYo = map[string]string{
	"saber": "se",
	"dar":   "doy",
	"ver":   "veo",
	"estar": "estoy",
	"caber": "quepo",
}

func getIrregularYoStem(verb string, base string) string {
	// -ger and -gir verbs
	if utils.HasOneOfMultipleSuffixes(verb, "ger", "gir") {
		return base[:len(base)-1] + "j"
	}

	// -guir verbs
	if strings.HasSuffix(verb, "eguir") {
		return base[:len(base)-3] + "ig"
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

	return getIrregularStem(verb, base)
}

func getIrregularStem(verb string, base string) string {
	// u -> uy
	if strings.HasSuffix(verb, "uir") {
		return utils.ChangeStem(base, "u", "uy")
	}

	// o -> oy
	if strings.HasSuffix(verb, "oir") {
		return utils.ChangeStem(base, "o", "oy")
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
		return getIrregularYoStem(verb, base)
	} else if subject == "nosotros" || subject == "vosotros" {
		return base
	} else {
		return getIrregularStem(verb, base)
	}
}

func GetPresentEnding(ending string, subject string) string {
	switch subject {
	case "yo":
		return "o"
	case "tu":
		if ending == "ar" {
			return "as"
		} else {
			return "es"
		}
	case "usted":
		if ending == "ar" {
			return "a"
		} else {
			return "e"
		}
	case "nosotros":
		if ending == "ar" {
			return "amos"
		} else if ending == "er" {
			return "emos"
		} else {
			return "imos"
		}
	case "vosotros":
		if ending == "ar" {
			return "ais"
		} else if ending == "er" {
			return "eis"
		} else {
			return "is"
		}
	case "ustedes":
		if ending == "ar" {
			return "an"
		} else {
			return "en"
		}
	default:
		return ""
	}
}

func ConjugatePresent(verb string, subject string) string {
	switch verb {
	case "ser":
		return ser[subject]
	case "ir":
		return ir[subject]
	case "haber":
		return haber[subject]
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
