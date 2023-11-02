package name

import "strings"

func Title(w string) string {
	runes := []rune(strings.ToLower(w))
	if len(runes) == 0 {
		return ""
	}
	if isLower(runes[0]) {
		runes[0] += 'A' - 'a'
	}
	return string(runes)
}
