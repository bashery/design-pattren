package luhn

import (
	"strconv"
	"strings"
)

// Valid checks is valid luhn code or not
func Valid(code string) bool {

	code = strings.ReplaceAll(code, " ", "")
	chars := strings.Split(code, "")

	// filter invalid lenght
	if len(chars) < 2 {
		return false
	}

	ind := len(chars) - 1
	rev := make([]int, len(chars))

	// convete item to int and reverse it
	for i := 0; i < len(chars); i++ {
		c, err := strconv.Atoi(chars[i])
		if err != nil {
			return false
		}
		rev[ind] = c
		ind--
	}

	// handle second item
	for k, v := range rev {
		if k%2 == 0 {
			continue
		}
		v *= 2
		if v > 9 {
			rev[k] = v - 9
			continue
		}
		rev[k] = v

	}

	// sum total of all items
	total := 0
	for _, n := range rev {
		total += n
	}

	// filter true result
	if total%10 == 0 {
		return true
	}

	return false
}
