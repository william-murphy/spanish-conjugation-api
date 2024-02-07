package indicative

// var preteriteSer = map[string]string{
// 	"yo":       "soy",
// 	"tu":       "eres",
// 	"usted":    "es",
// 	"nosotros": "somos",
// 	"vosotros": "soi",
// 	"ustedes":  "son",
// }

// var preteriteHaber = map[string]string{
// 	"yo":       "he",
// 	"tu":       "has",
// 	"usted":    "ha",
// 	"nosotros": "hemos",
// 	"vosotros": "habeis",
// 	"ustedes":  "han",
// }

// var preteriteIr = map[string]string{
// 	"yo":       "voy",
// 	"tu":       "vas",
// 	"usted":    "va",
// 	"nosotros": "vamos",
// 	"vosotros": "vais",
// 	"ustedes":  "van",
// }

// var preteriteArEndings = map[string]string{
// 	"yo":       "o",
// 	"tu":       "as",
// 	"usted":    "a",
// 	"nosotros": "amos",
// 	"vosotros": "áis",
// 	"ustedes":  "an",
// }

// var preteriteErIrEndings = map[string]string{
// 	"yo":       "o",
// 	"tu":       "es",
// 	"usted":    "e",
// 	"nosotros": "emos",
// 	"vosotros": "éis",
// 	"ustedes":  "en",
// }

// func getIrregularPreteriteYoStem(verb string, base string) string {
// 	// quir verbs
// 	if strings.HasSuffix(verb, "quir") {
// 		return utils.ChangeStem(base, "qu", "c")
// 	}

// 	// -ger and -gir verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "ger", "gir") {
// 		return base[:len(base)-1] + "j"
// 	}

// 	// -eguir verbs
// 	if strings.HasSuffix(verb, "eguir") {
// 		return base[:len(base)-3] + "ig"
// 	}

// 	// -guir verbs
// 	if strings.HasSuffix(verb, "guir") {
// 		return base[:len(base)-1]
// 	}

// 	// -igo verbs
// 	if strings.HasSuffix(verb, "decir") {
// 		return base[:len(base)-2] + "ig"
// 	}

// 	// -ago verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "hacer", "facer") {
// 		return base[:len(base)-1] + "g"
// 	}

// 	// -asgo verbs
// 	if strings.HasSuffix(verb, "asir") {
// 		return base + "g"
// 	}

// 	// -algo verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "salir", "valer") {
// 		return base + "g"
// 	}

// 	// -oir and -aer verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "oir", "aer") {
// 		return base + "ig"
// 	}

// 	// -engo and -ongo verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "tener", "poner") {
// 		return base + "g"
// 	}

// 	// consonant + -cer and -cir verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "cer", "cir") && !utils.ArrayContainsString(utils.Vowels, base[:len(base)-2]) {
// 		return base[:len(base)-1] + "z"
// 	}

// 	// vowel + -cer and -cir verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "cer", "cir") && utils.ArrayContainsString(utils.Vowels, base[:len(base)-2]) {
// 		return base[:len(base)-1] + "zc"
// 	}

// 	return getIrregularPreteriteStem(verb, base)
// }

// func getIrregularPreteriteStem(verb string, base string) string {
// 	// -uir verbs
// 	if utils.HasOneOfMultipleSuffixes(verb, "oir", "uir") {
// 		return base + "y"
// 	}

// 	// oler
// 	if verb == "oler" {
// 		return utils.ChangeStem(base, "o", "hue")
// 	}

// 	// e -> ie
// 	if utils.HasOneOfMultipleSuffixes(verb, "certar", "vesar", "cerrar", "fesar", "pertar", "pezar", "regar", "nevar", "mendar", "sentar", "lentar", "menzar", "helar", "bernar", "negar", "pensar", "querer", "fender", "perder", "tender", "cender", "tener", "mentir", "gerir", "vertir", "ferir", "sentir", "venir") {
// 		return utils.ChangeStem(base, "e", "ie")
// 	}

// 	// e -> i
// 	if utils.HasOneOfMultipleSuffixes(verb, "eguir", "decir", "elegir", "medir", "reir", "servir", "vestir", "regir", "pedir", "petir") {
// 		return utils.ChangeStem(base, "e", "i")
// 	}

// 	// o -> ue
// 	if utils.HasOneOfMultipleSuffixes(verb, "sonar", "morzar", "costar", "mostrar", "probar", "cordar", "ronar", "volar", "rogar", "contrar", "contar", "colgar", "solver", "cocer", "mover", "volver", "morder", "soler", "torcer", "poder", "moler", "llover", "doler", "dormir", "morir") {
// 		return utils.ChangeStem(base, "o", "ue")
// 	}

// 	// i -> ie
// 	if strings.HasSuffix(verb, "quirir") {
// 		return utils.ChangeStem(base, "i", "ie")
// 	}

// 	// u -> ue
// 	if strings.HasSuffix(verb, "jugar") {
// 		return utils.ChangeStem(base, "u", "ue")
// 	}

// 	return base
// }

// func GetPreteriteStem(verb string, base string, subject string) string {
// 	if subject == "yo" {
// 		return getIrregularPreteriteYoStem(verb, base)
// 	} else if subject == "nosotros" || subject == "vosotros" {
// 		return base
// 	} else {
// 		return getIrregularPreteriteStem(verb, base)
// 	}
// }

// func GetPreteriteEnding(ending string, subject string) string {
// 	if ending == "ar" {
// 		return preteriteArEndings[subject]
// 	} else {
// 		return preteriteErIrEndings[subject]
// 	}
// }

// func ConjugatePreterite(verb string, subject string) string {
// 	switch verb {
// 	case "ser":
// 		return preteriteSer[subject]
// 	case "ir":
// 		return preteriteIr[subject]
// 	case "haber":
// 		return preteriteHaber[subject]
// 	}

// 	value, exists := irregularYo[verb]
// 	if subject == "yo" && exists {
// 		return value
// 	}

// 	base, ending := utils.SplitVerb(verb)

// 	stem := GetPreteriteStem(verb, base, subject)
// 	newEnding := GetPreteriteEnding(ending, subject)

// 	return stem + newEnding
// }
