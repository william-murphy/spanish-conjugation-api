package conjugation

func Conjugate(verb string, mood string, tense string, subject string) string {
	if mood == "indicative" {
		return conjugateIndicative(verb, tense, subject)
	} else if mood == "subjunctive" {
		return conjugateSubjunctive(verb, tense, subject)
	} else {
		return conjugateImperative(verb, subject)
	}
}

func conjugateIndicative(verb string, tense string, subject string) string {
	return ""
}

func conjugateSubjunctive(verb string, tense string, subject string) string {
	return ""
}

func conjugateImperative(verb string, subject string) string {
	return ""
}
