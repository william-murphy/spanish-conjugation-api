package subjunctive

import "spanish-conjugation-api/api/conjugation/pastparticiple"

func ConjugatePluperfect(verb string, subject string) string {
	return ConjugateImperfect("haber", subject) + " " + pastparticiple.CreatePastParticiple(verb)
}
