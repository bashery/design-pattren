package grains

import "errors"

/*
1  2  3  4
1  2  4  8
*/

// this code is not good enoght. the mathimathic way is better.
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("invalid input")
	}
	var result uint64 = 1

	for i := 1; i < input; i++ {
		result = result * 2
	}
	return result, nil
}

func Total() uint64 {
	// the result is already known.
	// when no input argument i prefere return fix value
	return 18446744073709551615
}
