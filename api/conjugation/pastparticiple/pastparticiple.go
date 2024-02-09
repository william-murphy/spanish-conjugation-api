package pastparticiple

import (
	"spanish-conjugation-api/api/utils"
	"strings"
)

var irregulars = map[string]string{
	"brir":     "bierto",
	"morir":    "muerto",
	"poner":    "puesto",
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
			return verb[:len(verb)-len(key)] + value
		}
	}

	stem, ending := utils.SplitVerb(verb)

	if ending == "ar" {
		return stem + "ado"
	} else {
		if utils.HasOneOfMultipleSuffixes(stem, "a", "e", "o") {
			return stem + "ído"
		}
		return stem + "ido"
	}
}
