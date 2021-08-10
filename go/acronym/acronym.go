package acronym

import (
	"strings"
)

func Abbreviate(s string) string {

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	//s = strings.ReplaceAll(s, "'", " ")
	words := strings.Split(s, " ")
	slc := make([]string, 1)

	for _, word := range words {

		word = strings.TrimRight(word, "_ -")
		if word == "" {
			continue
		}
		word = strings.Title(word)

		slc = append(slc, string(word[0]))
	}

	acr := strings.Join(slc, "")
	return acr
}
