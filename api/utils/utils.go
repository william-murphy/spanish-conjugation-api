package utils

import (
	"regexp"
	"strings"
)

var Vowels = []string{"a", "e", "i", "o", "u"}

func ArrayContainsString(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

func IsIrregular(input map[string]string, target string) string {
	for key, value := range input {
		if strings.HasSuffix(target, key) {
			index := len(target) - len(key)
			beginning := target[:index]
			return beginning + value
		}
	}
	return ""
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
