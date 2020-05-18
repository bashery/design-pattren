package sublist

import (
	"reflect"
)

// Relation is "equal" or "unequal" or  "sublist" or "superlist"
type Relation string

func Sublist(slcA, slcB []int) Relation {
	lena := len(slcA)
	lenb := len(slcB)

	// find equal lists
	if reflect.DeepEqual(slcA, slcB) {
		return "equal"
	}

	// find sublist
	if lena < lenb {
		for i := 0; i < lenb; i++ {
			if reflect.DeepEqual(slcA, slcB[i:i+lena]) {
				return "sublist"
			}
			if i+lena == lenb {
				break
			}
		}
	}

	// find superlist
	if lena > lenb {
		for i := 0; i < lena; i++ {
			if reflect.DeepEqual(slcB, slcA[i:i+lenb]) {
				return "superlist"
			}
			if i+lenb == lena {
				break
			}
		}
	}
	return "unequal"
}
