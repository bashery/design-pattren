/*
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
======================================================================================
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	lastChar := remark[len(remark)-1]
	// 保证remark 不为空 且为字母
	if remark == strings.ToUpper(remark) && remark != strings.ToLower(remark) {
		if lastChar == '?' {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if lastChar == '?' {
		return "Sure."
	}
	return "Whatever."

}
func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
=====================================================================================
*/
package bob

import "strings"

// gives knowitall answer to strings
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	screamedRemark := strings.ToUpper(remark)
	isScreamed := remark == screamedRemark
	isJibberish := strings.ToUpper(remark) == strings.ToLower(remark)
	isQuestion := remark[len(remark)-1] == '?'

	if isQuestion {
		if isScreamed && !isJibberish {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if isScreamed && !isJibberish {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
