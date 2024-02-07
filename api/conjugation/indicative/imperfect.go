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

var imperfectArEndings = map[string]string{
	"yo":       "aba",
	"tu":       "abas",
	"usted":    "aba",
	"nosotros": "ábamos",
	"vosotros": "abais",
	"ustedes":  "aban",
}

var imperfectErIrEndings = map[string]string{
	"yo":       "ía",
	"tu":       "ías",
	"usted":    "ía",
	"nosotros": "íamos",
	"vosotros": "ías",
	"ustedes":  "ían",
}

func GetImperfectEnding(ending string, subject string) string {
	// TODO - add error handling to functions like these
	if ending == "ar" {
		return imperfectArEndings[subject]
	} else {
		return imperfectErIrEndings[subject]
	}
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
