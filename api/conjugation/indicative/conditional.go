package indicative

func ConjugateConditional(verb string, subject string) string {
	stem := GetFutureStem(verb)
	ending := GetImperfectEnding("er", subject)
	return stem + ending
}
