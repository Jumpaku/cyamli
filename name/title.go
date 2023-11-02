package name

import "strings"

func Title(w string) string {
	runes := []rune(strings.ToLower(w))
	if len(runes) == 0 {
		return ""
	}
	runes[0] += 'A' - 'a'
	return string(runes)
}
