package indicative

import "spanish-conjugation-api/api/conjugation/pastparticiple"

func ConjugateConditionalPerfect(verb string, subject string) string {
	return ConjugateConditional("haber", subject) + " " + pastparticiple.CreatePastParticiple(verb)
}
