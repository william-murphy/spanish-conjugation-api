package gerund

import "spanish-conjugation-api/api/utils"

func GetGerund(verb string) string {
	stem, ending := utils.SplitVerb(verb)

	if ending == "ar" {
		return stem + "ando"
	} else {
		return ""
	}
}
