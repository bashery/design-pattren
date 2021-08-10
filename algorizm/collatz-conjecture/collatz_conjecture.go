package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(i int) (int, error) {
	if i < 1 {
		return 0, fmt.Errorf("cant handle 0 or nigative number")
	}
	staps := 0
	for i > 1 {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = (i * 3) + 1
		}
		staps++
	}
	return staps, nil
}
