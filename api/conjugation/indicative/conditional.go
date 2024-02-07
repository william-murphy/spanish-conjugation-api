package indicative

var conditionalEndings = map[string]string{
	"yo":       "ía",
	"tu":       "ías",
	"usted":    "ía",
	"nosotros": "íamos",
	"vosotros": "íais",
	"ustedes":  "ían",
}

func ConjugateConditional(verb string, subject string) string {
	stem := GetFutureStem(verb)
	ending, exists := conditionalEndings[subject]
	if !exists {
		return "Invalid subject given."
	}
	return stem + ending
}
