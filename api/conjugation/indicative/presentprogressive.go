package indicative

import "spanish-conjugation-api/api/conjugation/gerund"

func ConjugatePresentProgressive(verb string, subject string) string {
	return ConjugatePresent("estar", subject) + " " + gerund.CreateGerund(verb)
}
