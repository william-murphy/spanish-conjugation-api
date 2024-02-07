package indicative

import "spanish-conjugation-api/api/conjugation/gerund"

func ConjugateConditionalProgressive(verb string, subject string) string {
	return ConjugateConditional("estar", subject) + " " + gerund.CreateGerund(verb)
}
