package name

import (
	"strings"

	"github.com/samber/lo"
)

type Path []string

func MakePath(s string) Path {
	runes := []rune{}
	for _, r := range strings.ToLower(s) {
		switch {
		case isUpper(r), isLower(r), isDigit(r):
			runes = append(runes, r)
		default:
			if len(runes) > 0 && runes[len(runes)-1] != ' ' {
				runes = append(runes, ' ')
			}
		}
	}
	if len(runes) == 0 {
		return Path{}
	}
	return strings.Split(string(runes), " ")
}

func (name Path) Join(separator, prefix, suffix string) string {
	return prefix + strings.Join(name, separator) + suffix
}
func (name Path) Append(w string) Path {
	return append(append([]string{}, name...), w)
}
func (name Path) Map(f func(string) string) Path {
	return lo.Map(name, func(s string, i int) string { return f(s) })
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
