package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("no")
	}
	var diff int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}

	}
	return diff, nil
}
