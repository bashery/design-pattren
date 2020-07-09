package anagram

import (
	//	"fmt"
	"sort"
	"strings"
)

// Detect resurn word list that containt sam word input letteral
func Detect(str string, strs []string) []string {
	str = strings.ToLower(str)
	sbj := sortStr(str)
	res := make([]string, 0)

	for _, w := range strs {
		match := w
		w = strings.ToLower(w)
		if str == w {
			continue
		}
		word := sortStr(w)
		if sbj == word {
			res = append(res, match)
		}
	}

	return res
}

func sortStr(str string) string {
	slc := make([]int, 0)
	for _, c := range str {
		slc = append(slc, int(c))
	}
	sort.Ints(slc)
	res := ""
	for _, c := range slc {
		res += string(int32(c))
	}
	return res
}
