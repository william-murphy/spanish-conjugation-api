package indicative

import "spanish-conjugation-api/api/conjugation/gerund"

func ConjugateFutureProgressive(verb string, subject string) string {
	return ConjugateFuture("estar", subject) + " " + gerund.CreateGerund(verb)
}
