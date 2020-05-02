package bob

import (
	"strings"
	"unicode"
)

func hasNoLetters(str string) bool {
	for _, char := range str {
		if unicode.IsLetter(char) {
			return false
		}
	}

	return true
}

func isUpper(str string) bool {
	for _, char := range str {
		if strings.ToUpper(string(char)) != string(char) {
			return false
		}
	}

	return true
}

func isQuestion(str string) bool {
	return len(str) > 0 && string(str[len(str)-1]) == "?"
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmedRemark := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(remark, "\t", ""), "\n", ""), " ", ""), "\r", "")

	switch {
	case len(trimmedRemark) == 0:
		return "Fine. Be that way!"
	case hasNoLetters(trimmedRemark) && !isQuestion(trimmedRemark):
		return "Whatever."
	case !hasNoLetters(trimmedRemark) && isUpper(trimmedRemark) && isQuestion(trimmedRemark):
		return "Calm down, I know what I'm doing!"
	case !hasNoLetters(trimmedRemark) && isUpper(trimmedRemark):
		return "Whoa, chill out!"
	case isQuestion(trimmedRemark):
		return "Sure."
	default:
		return "Whatever."
	}
}
