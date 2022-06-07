package chgk

import "strings"

func maskString(s string) string {
	var sb strings.Builder
	for s := range []rune(strings.ToLower(s)) {
		if s >= 'a' && s <= 'z' {
			sb.WriteRune('x')
		} else if s >= 'A' && s <= 'Z' {
			sb.WriteRune('X')
		} else {
			sb.WriteRune(rune(s))
		}
	}
	return sb.String()
}
