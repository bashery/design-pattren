package accumulate

import (
//"fmt"
)

func Accumulate(arr []string, f func(s string) string) (result []string) {
	for _, val := range arr {
		result = append(result, f(val))
	}
	return result
}
