package reverse

func Reverse(s string) string {
	r := []rune(s)
	l := len(r) - 1
	res := ""
	for i := l; i >= 0; i-- {
		res = res + string(r[i])
	}
	return res
}
