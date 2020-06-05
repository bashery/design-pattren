package anagram

import "fmt"

func Detect(str string, slc []string) []string {
	sublist := []rune{}
	for _, v := range slc {
		for _, c := range v {
			sublist = append(sublist, c)

		}
	}
	fmt.Println(str, slc)
	return []string{}
}
