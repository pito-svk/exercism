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
	formattedRemark := strings.ReplaceAll(remark, " ", "")
	formattedRemark = strings.ReplaceAll(formattedRemark, "\t", "")
	formattedRemark = strings.ReplaceAll(formattedRemark, "\r", "")
	formattedRemark = strings.ReplaceAll(formattedRemark, "\n", "")

	switch {
	case len(formattedRemark) == 0:
		return "Fine. Be that way!"
	case hasNoLetters(formattedRemark) && !isQuestion(formattedRemark):
		return "Whatever."
	case !hasNoLetters(formattedRemark) && isUpper(formattedRemark) && isQuestion(formattedRemark):
		return "Calm down, I know what I'm doing!"
	case !hasNoLetters(formattedRemark) && isUpper(formattedRemark):
		return "Whoa, chill out!"
	case isQuestion(formattedRemark):
		return "Sure."
	default:
		return "Whatever."
	}
}
