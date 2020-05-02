package acronym

import (
	"strings"
)

// Abbreviate creates acronym from string
func Abbreviate(s string) string {
	words := strings.Split(s, " ")

	var acronym string

	for _, word := range words {
		word = strings.Trim(word, "_")

		if strings.Contains(word, "-") && word != "-" {
			upperWord := strings.ToUpper(word)

			firstLetter := string(strings.Split(upperWord, "-")[0][0])
			secondLetter := string(strings.Split(upperWord, "-")[1][0])

			acronym += firstLetter + secondLetter
		} else if word != "-" {
			acronym += strings.ToUpper(string(word[0]))
		}
	}

	return acronym
}
