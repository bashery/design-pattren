package series

func All(l int, s string) (slc []string) {
	ln := len(s) - l
	for i := 0; i <= ln; i++ {
		slc = append(slc, s[i:i+l])
	}

	return slc
}

func UnsafeFirst(i int, s string) string {
	return string(s[:i])
}

func First(i int, s string) (string, bool) {
	if len(s) < 0 {
		return "", false
	}
	return string(s[:i]), true
}
