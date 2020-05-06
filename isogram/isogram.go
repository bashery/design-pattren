package isogram

import (
	"strings"
)

func IsIsogram(str string) bool {
	word := strings.ToLower(str)
	//word = strings.ReplaceAll(word, "-", "")
	//word = strings.ReplaceAll(word, " ", "")
	l := len(word)
	if l <= 1 {
		return true
	}
	for i := 0; i < l; i++ {
		for c := i + 1; c < l; c++ {
			// for benchmarck I complicate this if statment rather then init strings
			if word[i] == word[c] && string(word[i]) != " " && string(word[i]) != "-" {
				return false
			}

		}
	}
	return true
}
