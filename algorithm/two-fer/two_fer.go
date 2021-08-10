package twofer

import "strings"

func ShareWith(name string) string {
	message := "One for you, one for me."
	if name == "" {
		return message
	}
	return strings.Replace(message, "you", name, 1)
}
