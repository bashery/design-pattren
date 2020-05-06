package raindrops

import (
	"strconv"
)

func Convert(num int) string {
	result := ""

	for i := 1; i < num; i++ {
		if i*3 == num {
			result = "Pling" + result
			continue
		}
		if i*5 == num {
			result = "Plang" + result
			continue
		}
		if i*7 == num {
			result = "Plong" + result
			continue
		}
	}

	if result == "" {
		result = strconv.Itoa(num)
	}

	return result
}
