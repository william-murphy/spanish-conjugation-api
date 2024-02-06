package utils

import (
	"regexp"
	"strings"
)

var Vowels = []string{"a", "e", "i", "o", "u"}

func ChangeStem(base, toReplace, replaceWith string) string {
	lastIndex := strings.LastIndex(base, toReplace)

	if lastIndex == -1 {
		return base
	}

	result := base[:lastIndex] + replaceWith + base[lastIndex+len(toReplace):]
	return result
}

func HasOneOfMultipleSuffixes(base string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		baseLen := len(base)
		suffixLen := len(suffix)
		if baseLen >= suffixLen && base[baseLen-suffixLen:] == suffix {
			return true
		}
	}
	return false
}

func ArrayContainsString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

func ContainsOnlyLowerCase(s string) bool {
	pattern := "^[a-z]+$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

func VerifyVerbEnding(verb string) bool {
	return strings.HasSuffix(verb, "ar") || strings.HasSuffix(verb, "er") || strings.HasSuffix(verb, "ir")
}

func SplitVerb(verb string) (string, string) {
	index := len(verb) - 2
	return verb[:index], verb[index:]
}
