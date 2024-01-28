package pastparticiple

import "spanish-conjugation-api/api/utils"

var irregulars = map[string]string{
	"abrir":    "abierto",
	"morir":    "muerto",
	"poner":    "puesto",
	"cubrir":   "cubierto",
	"decir":    "dicho",
	"romper":   "roto",
	"escribir": "escrito",
	"efacer":   "efacto",
	"facer":    "fecho",
	"hacer":    "hecho",
	"solver":   "suelto",
	"volver":   "vuelto",
	"ver":      "visto",
}

func GetPastParticiple(verb string) string {
	irreg := utils.IsIrregular(irregulars, verb)

	if irreg != "" {
		return irreg
	}

	stem, ending := utils.SplitVerb(verb)

	if ending == "ar" {
		return stem + "ado"
	} else {
		return stem + "ido"
	}
}
