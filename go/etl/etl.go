// convert data from  1:["a", "b"]} to {"a":1, "b":1}
package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	result := map[string]int{}
	for k, listv := range input {
		for _, item := range listv {
			result[strings.ToLower(item)] = k
		}
	}
	return result
}
