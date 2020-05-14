package pangram

import (
	"strings"
)

func IsPangram(str string) bool {
	str = strings.ToLower(str)
	chars := 0

	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		for _, v := range str {
			if c == v {
				chars++
				break
			}
		}
	}
	return chars == 26
}
