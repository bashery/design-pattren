package wordcount

import (
	"strings"
)

type Frequency map[string]int // Using this return type.

func WordCount(phrase string) Frequency {

	frequency := Frequency{}

	// some init
	phrase = strings.ToLower(phrase)
	phrase = strings.ReplaceAll(phrase, ",", " ")
	phrase = strings.ReplaceAll(phrase, "!!&@$%^&", "")
	phrase = strings.ReplaceAll(phrase, "\n", " ")

	slc := strings.Split(phrase, " ")

	for _, w := range slc {
		if w == "" {
			continue
		}
		if w[len(w)-1] == ':' || w[len(w)-1] == '.' {
			w = w[:len(w)-1]
		}
		if string(w[len(w)-1]) == "'" {
			w = strings.Trim(w, "'")
		}

		frequency[w]++
	}
	return frequency

}
