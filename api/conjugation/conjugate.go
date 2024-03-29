package conjugation

import (
	"spanish-conjugation-api/api/conjugation/imperative"
	"spanish-conjugation-api/api/conjugation/indicative"
	"spanish-conjugation-api/api/conjugation/subjunctive"
)

func Conjugate(verb string, mood string, tense string, subject string) string {
	// TODO - decision trees?? lots of if else statements
	if mood == "indicative" {
		return conjugateIndicative(verb, tense, subject)
	} else if mood == "subjunctive" {
		return conjugateSubjunctive(verb, tense, subject)
	} else {
		return conjugateImperative(verb, subject)
	}
}

func conjugateIndicative(verb string, tense string, subject string) string {
	switch tense {
	case "present":
		return indicative.ConjugatePresent(verb, subject)
	case "presentprogressive":
		return indicative.ConjugatePresentProgressive(verb, subject)
	case "nearfuture":
		return indicative.ConjugateNearFuture(verb, subject)
	case "presentperfect":
		return indicative.ConjugatePresentPerfect(verb, subject)
	case "imperfect":
		return indicative.ConjugateImperfect(verb, subject)
	case "pluperfect":
		return indicative.ConjugatePluperfect(verb, subject)
	case "pastprogressive":
		return indicative.ConjugatePastProgressive(verb, subject)
	case "future":
		return indicative.ConjugateFuture(verb, subject)
	case "futureperfect":
		return indicative.ConjugateFuturePerfect(verb, subject)
	case "futureprogressive":
		return indicative.ConjugateFutureProgressive(verb, subject)
	case "conditional":
		return indicative.ConjugateConditional(verb, subject)
	case "conditionalperfect":
		return indicative.ConjugateConditionalPerfect(verb, subject)
	case "conditionalprogressive":
		return indicative.ConjugateConditionalProgressive(verb, subject)
	case "preterite":
		return indicative.ConjugatePreterite(verb, subject)
	}

	return "Invalid tense given for indicative mood."
}

func conjugateSubjunctive(verb string, tense string, subject string) string {
	switch tense {
	case "present":
		return subjunctive.ConjugatePresent(verb, subject)
	case "presentperfect":
		return subjunctive.ConjugatePresentPerfect(verb, subject)
	case "imperfect":
		return subjunctive.ConjugateImperfect(verb, subject)
	case "pluperfect":
		return subjunctive.ConjugatePluperfect(verb, subject)
	}

	return "Invalid tense given for subjunctive mood."
}

func conjugateImperative(verb string, subject string) string {
	switch subject {
	case "tu":
		return imperative.ConjugateTu(verb)
	case "vosotros":
		return imperative.ConjugateVosotros(verb)
	case "usted":
		return imperative.ConjugateUsted(verb)
	case "ustedes":
		return imperative.ConjugateUstedes(verb)
	case "nosotros":
		return imperative.ConjugateNosotros(verb)
	}

	return "Invalid subject given for imperative mood."
}
