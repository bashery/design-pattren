package phonenumber

import (
	"errors"
)

func Number(n string) (string, error) {
	digit := "0123456789"
	num := ""
	for _, n := range n {

		for _, d := range digit {
			if n == d {
				num += string(d)
			}
		}
	}
	if len(num) < 10 {
		return num, errors.New("err")
	}
	one := byte(49)
	if len(num) == 11 {
		if string(num[0]) == "1" && num[1] > one && num[4] > one {
			return num[1:], nil
		}
	}
	if len(num) == 10 {
		if num[0] > one && num[3] > one {
			return num, nil
		}
	}
	return n, errors.New("error")
}

func AreaCode(n string) (string, error) {
	num, err := Number(n)
	if err != nil {
		return num, err
	}
	return num[:3], nil
}

func Format(n string) (string, error) {
	num, err := Number(n)
	if err != nil {
		return num, err
	}
	return "(" + num[:3] + ") " + num[3:6] + "-" + num[6:], nil
}

/*
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Number(s string) (string, error) {
	builder := strings.Builder{}
	c, _ := regexp.Compile("^(\\+?1.?)?((\\(([2-9]\\d\\d)\\))|([2-9]\\d\\d))[^\\d]*([2-9]\\d\\d)[^\\d]*(\\d\\d\\d\\d)[^\\d]*$")
	matches := c.FindStringSubmatch(s)
	if len(matches) < 8 {
		return "", errors.New("error")
	}
	if matches[4] != "" {
		builder.WriteString(matches[4])
	} else {
		builder.WriteString(matches[5])
	}
	builder.WriteString(matches[6])
	builder.WriteString(matches[7])
	return builder.String(), nil
}

func AreaCode(s string) (string, error) {
	res, err := Number(s)
	if err != nil {
		return "", err
	}
	return res[0:3], nil
}

func Format(s string) (string, error) {
	res, err := Number(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", res[0:3], res[3:6], res[6:]), nil
}

*/
