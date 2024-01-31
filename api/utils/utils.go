package utils

import (
	"regexp"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

func GetVowels(exclude ...string) []string {
	// Create a map to store elements to be removed
	elementsToRemoveMap := make(map[string]struct{})
	for _, element := range exclude {
		elementsToRemoveMap[element] = struct{}{}
	}

	// Create a new slice to store the elements not in elementsToRemoveMap
	result := make([]string, 0, len(vowels))
	for _, element := range vowels {
		if _, exists := elementsToRemoveMap[element]; !exists {
			result = append(result, element)
		}
	}

	return result
}

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
