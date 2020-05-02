package bob

import (
	"strings"
	"unicode"
)

func hasLetters(str string) bool {
	for _, char := range str {
		if unicode.IsLetter(char) {
			return true
		}
	}

	return false
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

func isEmpty(str string) bool {
	return len(str) == 0
}

// Hey responds as Bob to Alice requests
func Hey(remark string) string {
	formattedRemark := strings.ReplaceAll(remark, " ", "")
	formattedRemark = strings.ReplaceAll(formattedRemark, "\t", "")
	formattedRemark = strings.ReplaceAll(formattedRemark, "\r", "")
	formattedRemark = strings.ReplaceAll(formattedRemark, "\n", "")

	if isEmpty(formattedRemark) {
		return "Fine. Be that way!"
	}

	if hasLetters(formattedRemark) && isUpper(formattedRemark) && isQuestion(formattedRemark) {
		return "Calm down, I know what I'm doing!"
	}

	if hasLetters(formattedRemark) && isUpper(formattedRemark) {
		return "Whoa, chill out!"
	}

	if isQuestion(formattedRemark) {
		return "Sure."
	}

	return "Whatever."
}
