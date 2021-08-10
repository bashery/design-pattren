package lsproduct

import (
	"errors"
)

// this take a string as number wth and return>>>
func LargestSeriesProduct(str string, l int) (int, error) {
	if l == 0 {
		return 1, nil
	}
	if l < 0 {
		return 1, errors.New("len str error!")
	}
	_, err := filter(str, l)
	if err != nil {
		return 0, err
	}
	if len(str) == l {
		return getEq(str, l)
	}

	total := 0
	for i := 0; i < len(str)-l; i++ {
		n := 1

		for j := i + 1; j <= i+l; j++ {
			n *= (int(str[j]) - 48)
		}
		if total < n {
			total = n
		}
	}

	return total, nil
}

func getEq(str string, l int) (int, error) {

	total := 0
	n := 1
	for i := 0; i < len(str); i++ {
		n *= (int(str[i]) - 48)
		if total < n {
			total = n
		}
	}

	return total, nil
}

// check if len and all tocke is str digit and return it
func filter(s string, l int) (string, error) {
	if len(s) < l {
		return "", errors.New("len str error!")
	}

	tokn := "0123456789"
	str := ""
	for _, c := range s {
		for _, t := range tokn {
			if c == t {
				str += string(t)
				break
			}
		}
	}
	if len(s) > len(str) {
		return "", errors.New("token error!")
	}
	return str, nil
}
