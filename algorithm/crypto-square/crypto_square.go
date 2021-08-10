package cryptosquare

import (
	"math"
	"strings"
)

func Encode(str string) string {
	if str == "" {
		return ""
	}
	str = filter(str)
	str = strings.ToLower(str)

	// the number of rows and columns
	r := int(math.Sqrt(float64(len(str))))
	c := int(len(str) / r)
	if r*c < len(str) {
		c += 1
	}

	for c-r > 1 {
		r += 1
		c -= 1
	}

	for i := len(str); i < r*c; i++ {
		str += " "
	}

	res := ""
	for col := 0; col < c; col++ {
		for row := 0; row < r; row++ {
			res += string(str[row*c+col])
		}
		res += " "
	}
	return res[:len(res)-1]
}

func filter(str string) string {
	pure := ""
	token := " !@#$%^&*.,"
	match := false
	for _, c := range str {
		for _, t := range token {
			if c == t {
				match = true
				break
			}
		}
		if !match {
			pure += string(c)
		}
		match = false
	}
	return pure
}
