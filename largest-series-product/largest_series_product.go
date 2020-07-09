package lsproduct

import (
	"errors"

	"fmt"
	"strconv"
)

// this take a string as number wth and return>>>
func LargestSeriesProduct(s string, l int) (int, error) {
	fmt.Println("input s is : ", s, "l is : ", l, "len of input is :", len(s))
	fmt.Println()
	s, err := check(s)
	if err != nil {
		return 0, err
	}
	if len(s) == l && l != 0 {
		return total1(s)
	}
	if len(s) > l {
		return total2(s, l)
	}
	if len(s) < l {
		return 1, errors.New("error")
	}

	return 1, nil //errors.New("error")
}

func total2(s string, l int) (int, error) {
	//list := make([]int, l)
	if l > len(s) {
		l = len(s) - 1
	}
	if l == 0 {
		return 1, nil
	}
	res := 0
	total := 0
	for i := 0; i < len(s)-l; i++ {
		n := int(s[i+1]) - 48

		if n == 0 {
			continue
		}

		total = n
		fmt.Println("n is ", n)
		list := make([]int, l)
		inc := 0

		for j := i + 1; j < i+l; j++ {

			total *= (int(s[j+1]) - 48)

			list[inc] = int(s[j+1]) - 48
			inc++
			fmt.Print("inc is : ", inc, "list is :", list)
		}
		//fmt.Println("\ntotal is :", total)
		fmt.Println()
		if total > res {
			res = total
		}
	}
	return res, nil
}

func total1(s string) (int, error) {
	list := make([]int, len(s))
	for i, v := range s {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, err
		}
		list[i] = n

	}
	res := list[0]
	for _, v := range list[1:] {
		res *= v
	}
	return res, nil

}
func check(s string) (string, error) {
	allow := "0123456789"
	res := ""
	for _, v := range allow {
		for _, c := range s {
			if v == c {
				res += string(c)
			}
		}
	}
	if len(res) != len(s) {
		return "", errors.New("error!  invalid string integer")
	}
	return res, nil
}
