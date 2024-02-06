package indicative

func ConjugateNearFuture(verb string, subject string) string {
	return ConjugatePresent("ir", subject) + " a " + verb
}
