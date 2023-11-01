package schema

import "strings"

type Name []string

func MakeName(s string) Name {
	runes := []rune{}
	for _, r := range strings.ToLower(s) {
		switch {
		case isUpper(r):
			if len(runes) > 0 && !isUpper(runes[len(runes)-1]) {
				runes = append(runes, ' ')
			}
			runes = append(runes, r-('A'-'a'))
		case isLower(r):
			if len(runes) > 0 && !isLower(runes[len(runes)-1]) {
				runes = append(runes, ' ')
			}
			runes = append(runes, r)
		case isDigit(r):
			if len(runes) > 0 && !isDigit(runes[len(runes)-1]) {
				runes = append(runes, ' ')
			}
			runes = append(runes, r)
		default:
			if len(runes) > 0 && runes[len(runes)-1] != ' ' {
				runes = append(runes, ' ')
			}
		}
	}
	return strings.Split(string(runes), " ")
}

func (name Name) Join(separator, prefix, suffix string) string {
	return prefix + strings.Join(name, separator) + suffix
}
func (name Name) Append(w string) Name {
	return append([]string(name), w)
}
func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}
func isLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
