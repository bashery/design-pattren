package bob

import "regexp"
import "strings"

func Hey(remark string) string {
	remark = strings.Trim(remark, " ")
	q := regexp.MustCompile(`.*\n*\?|\?\s*$`).MatchString(remark)
	yill := regexp.MustCompile(`[a-z]+`).MatchString(remark)
	yq := regexp.MustCompile(`[a-z0-9]+`).MatchString(remark)
	d := regexp.MustCompile(`[A-Za-z]+`).MatchString(remark)
	imj := regexp.MustCompile(`[a-zA-Z0-9]+`).MatchString(remark)

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case silent(remark):
		return "Fine. Be that way!"
	case q:
		if !yq {
			if !imj {
				return "Sure."
			}
			return "Calm down, I know what I'm doing!"
		}
		if string(remark[len(remark)-1]) != "?" {
			return "Whatever."
		}
		return "Sure."
	case !yill:
		if !d {
			return "Whatever."
		}
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func silent(srt string) bool {
	var c string

	for i := 0; i < len(srt); i++ {
		c = string(srt[i])
		if c == "\n" || c == "	" || c == "\r" || c == " " {
			// ok
		} else {
			return false
		}
	}
	return true
}
