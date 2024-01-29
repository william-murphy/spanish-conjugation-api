package gerund

import (
	"spanish-conjugation-api/api/utils"
	"strings"
)

func gerundStem(base string, ending string) string {
	if ending == "ar" {
		return base
	} else if ending == "er" {
		if strings.HasSuffix(base, "od") {
			return base[:len(base)-2] + "ud"
		} else {
			return base
		}
	} else {
		if strings.HasSuffix(base, "ent") {
			return base[:len(base)-3] + "int"
		} else if strings.HasSuffix(base, "ec") {
			return base[:len(base)-2] + "ic"
		} else if strings.HasSuffix(base, "orm") {
			return base[:len(base)-3] + "urm"
		} else if strings.HasSuffix(base, "ed") {
			return base[:len(base)-2] + "id"
		} else if strings.HasSuffix(base, "est") {
			return base[:len(base)-3] + "ist"
		} else if strings.HasSuffix(base, "or") {
			return base[:len(base)-2] + "ur"
		} else {
			return base
		}
	}

}

func gerundEnding(stem string, ending string) string {
	if ending == "ar" {
		return "ando"
	} else {
		if stem == "" || utils.ArrayContainsString(utils.Vowels, stem[len(stem)-1:]) {
			return "yendo"
		} else {
			return "iendo"
		}
	}
}

func CreateGerund(verb string) string {
	base, ending := utils.SplitVerb(verb)
	stem := gerundStem(base, ending)
	newEnding := gerundEnding(stem, ending)
	return stem + newEnding
}
