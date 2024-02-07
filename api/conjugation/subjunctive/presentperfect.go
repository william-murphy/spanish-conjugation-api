package subjunctive

import "spanish-conjugation-api/api/conjugation/pastparticiple"

func ConjugatePresentPerfect(verb string, subject string) string {
	return ConjugatePresent("haber", subject) + " " + pastparticiple.CreatePastParticiple(verb)
}
