package brackets

func Bracket(str string) bool {
	bracks := []string{"{}", "[]", "()"}
	openBracks := []string{"(", "{", "["}
	closeBacks := []string{"]", "}", ")"}
	opens := ""
	for _, r := range str {
		r := string(r)
		if in(r, openBracks) {
			opens += r
		}
		if in(r, closeBacks) {

			if len(opens) == 0 {
				return false
			}
			brc := string(opens[len(opens)-1]) + string(r)
			if match(brc, bracks) {

				opens = opens[:len(opens)-1]
			} else {
				return false
			}

		}
	}

	return len(opens) == 0
}

// match ckech if input is in right list brackit
func match(brc string, bracks []string) bool {
	for _, v := range bracks {
		if brc == v {
			return true
		}
	}
	return false
}

func in(c string, str []string) bool {
	for _, v := range str {
		if c == v {
			return true
		}
	}
	return false
}
