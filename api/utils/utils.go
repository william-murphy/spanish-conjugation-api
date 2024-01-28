package utils

import (
	"regexp"
	"strings"
)

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
