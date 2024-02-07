package indicative

import "spanish-conjugation-api/api/conjugation/gerund"

func ConjugatePastProgressive(verb string, subject string) string {
	return ConjugateImperfect("estar", subject) + " " + gerund.CreateGerund(verb)
}
