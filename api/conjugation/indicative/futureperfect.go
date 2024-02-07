package indicative

import "spanish-conjugation-api/api/conjugation/pastparticiple"

func ConjugateFuturePerfect(verb string, subject string) string {
	return ConjugateFuture("haber", subject) + " " + pastparticiple.CreatePastParticiple(verb)
}
