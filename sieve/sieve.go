package sieve

import (
	"fmt"
)

func Sieve(up int) []int {
	fmt.Println("new test with : ", up)
	res := []int{}
	//if up == 2 {		res = append(res, 2)	}

	for i := 2; i <= up; i++ {
		prime := true
		for j := 2; j <= i; j++ {

			fmt.Println(i, "%", j, "=", i%j)
			if j == i {
				break
			}
			if i%j == 0 {
				fmt.Println()

				prime = false
				continue
			}

		}
		if prime == true {

			fmt.Println(i, " ok")
			res = append(res, i)
		}

	}
	fmt.Println(res)
	return res
}
