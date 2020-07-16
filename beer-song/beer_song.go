package beer

import (
	"fmt"
)

func Verses(a, b int) (string, error) {
	if a < b || b < 0 || a > 99 {
		return "", fmt.Errorf("an error")
	}
	res := ""
	for i := a; i >= b; i-- {
		s, _ := Verse(i)
		res += s + "\n"
	}
	return res, nil
}

func Song() string {
	res, _ := Verses(99, 0)
	return res
}

func Verse(a int) (string, error) {
	s := "s"
	if a == 2 {
		s = ""
	}
	if a == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if a == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if a > 99 {
		return "", fmt.Errorf("an error")
	}

	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle%s of beer on the wall.\n", a, a, a-1, s), nil
}
