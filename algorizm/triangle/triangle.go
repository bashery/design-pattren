package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	//var k Kind
	var slc = []float64{a, b, c}
	var inf = math.Inf(1)
	if a+b < c || a+c < b || b+c < a || a == inf || b == inf || c == inf {
		return NaT
	}
	for _, v := range slc {
		if math.IsNaN(v) {
			return NaT
		}
		if v <= 0 {
			return NaT
		}
	}

	switch {
	case a == b && a == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	case b == c && b != a:
		return Iso
	}
	return Sca
}
