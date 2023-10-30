package cliautor

import "strings"

func Words(s string) []string {
	runes := []rune{}
	for _, r := range strings.ToLower(s) {
		switch {
		case IsUpper(r):
			if len(runes) > 0 && !IsUpper(runes[len(runes)-1]) {
				runes = append(runes, ' ')
			}
			runes = append(runes, r-('A'-'a'))
		case IsLower(r):
			if len(runes) > 0 && !IsLower(runes[len(runes)-1]) {
				runes = append(runes, ' ')
			}
			runes = append(runes, r)
		case IsDigit(r):
			if len(runes) > 0 && !IsDigit(runes[len(runes)-1]) {
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
func IsUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}
func IsLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}
func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
