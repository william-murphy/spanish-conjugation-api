package gerund

import (
	"spanish-conjugation-api/api/utils"
)

func getGerundStem(verb string, base string) string {
	// o -> u
	if utils.HasOneOfMultipleSuffixes(verb, "oder", "ormir", "orir") {
		return utils.ChangeStem(base, "o", "u")
	}

	// e -> i
	if utils.HasOneOfMultipleSuffixes(verb, "entir", "ecir", "edir", "estir") {
		return utils.ChangeStem(base, "e", "i")
	}

	return base
}

func getGerundEnding(base string, ending string) string {
	if ending == "ar" {
		return "ando"
	} else {
		if base == "" || utils.ArrayContainsString(utils.Vowels, base[len(base)-1:]) {
			return "yendo"
		} else {
			return "iendo"
		}
	}
}

func CreateGerund(verb string) string {
	base, ending := utils.SplitVerb(verb)
	stem := getGerundStem(verb, base)
	newEnding := getGerundEnding(stem, ending)
	return stem + newEnding
}
