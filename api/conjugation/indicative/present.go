package indicative

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

var fullyIrregularYoForm = map[string]string{
	"saber": "se",
	"dar":   "doy",
	"ver":   "veo",
	"oir":   "oigo",
	"estar": "estoy",
}

var irregularYoFormOnly = map[string]string{
	"guir":    "g",
	"ger":     "j",
	"gir":     "j",
	"caber":   "quep",
	"caer":    "caig",
	"conocer": "conozc",
	"hacer":   "hag",
	"poner":   "pong",
	"salir":   "salg",
	"ucir":    "uzc",
	"traer":   "traig",
	"valer":   "valg",
	"uir":     "uy",
}

func ConjugatePresentIndicative(verb string, subject string) string {
	return ""
	// irreg := utils.IsIrregular(irregulars, verb)

	// if irreg != "" {
	// 	return irreg
	// }

	// stem, ending := utils.SplitVerb(verb)

	// if ending == "ar" {
	// 	return stem + "ado"
	// } else {
	// 	return stem + "ido"
	// }
}
