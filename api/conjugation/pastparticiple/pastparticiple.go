package pastparticiple

import (
	"spanish-conjugation-api/api/utils"
	"strings"
)

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
}

func CreatePastParticiple(verb string) string {

	// Check if ver
	if verb == "ver" {
		return "visto"
	}

	// Check if irregular
	for key, value := range irregulars {
		if strings.HasSuffix(verb, key) {
			index := len(verb) - len(key)
			beginning := verb[:index]
			return beginning + value
		}
	}

	stem, ending := utils.SplitVerb(verb)

	if ending == "ar" {
		return stem + "ado"
	} else {
		return stem + "ido"
	}
}
