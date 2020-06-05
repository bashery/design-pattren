package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// insubmet solution
func Valid(code string) bool {
	fmt.Println(code)
	code = strings.ReplaceAll(code, " ", "")
	if len(code) < 2 {
		return false
	}

	slc := []int{}
	n := 0
	for _, v := range code {
		n, _ = strconv.Atoi(string(v))
		slc = append(slc, n)
	}

	sum := 0

	for i := len(slc) - 1; i >= 0; i-- {
		if len(slc)%2 == 0 {
			slc[i] = slc[i] * 2
			if slc[i] > 9 {
				slc[i] = slc[i] - 9
			}

		}
		sum += slc[i]

	}
	if sum%10 == 0 {
		return true
	}
	return false
}
