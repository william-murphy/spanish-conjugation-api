package indicative

import (
	"spanish-conjugation-api/api/utils"
)

var imperfectSer = map[string]string{
	"yo":       "era",
	"tu":       "eras",
	"usted":    "era",
	"nosotros": "éramos",
	"vosotros": "erais",
	"ustedes":  "eran",
}

var imperfectIr = map[string]string{
	"yo":       "iba",
	"tu":       "ibas",
	"usted":    "iba",
	"nosotros": "íbamos",
	"vosotros": "ibais",
	"ustedes":  "iban",
}

var imperfectVer = map[string]string{
	"yo":       "veía",
	"tu":       "veías",
	"usted":    "veía",
	"nosotros": "veíamos",
	"vosotros": "veías",
	"ustedes":  "veían",
}

func GetImperfectEnding(ending string, subject string) string {
	// TODO - add error handling to functions like these

	switch subject {
	case "yo":
		if ending == "ar" {
			return "aba"
		} else {
			return "ía"
		}
	case "tu":
		if ending == "ar" {
			return "abas"
		} else {
			return "ías"
		}
	case "usted":
		if ending == "ar" {
			return "aba"
		} else {
			return "ía"
		}
	case "nosotros":
		if ending == "ar" {
			return "ábamos"
		} else {
			return "íamos"
		}
	case "vosotros":
		if ending == "ar" {
			return "abais"
		} else {
			return "íais"
		}
	case "ustedes":
		if ending == "ar" {
			return "aban"
		} else {
			return "ían"
		}
	}
	return "Invalid verb given."
}

func ConjugateImperfect(verb string, subject string) string {
	switch verb {
	case "ser":
		return imperfectSer[subject]
	case "ir":
		return imperfectIr[subject]
	case "ver":
		return imperfectVer[subject]
	}

	base, ending := utils.SplitVerb(verb)
	newEnding := GetImperfectEnding(ending, subject)

	return base + newEnding
}
