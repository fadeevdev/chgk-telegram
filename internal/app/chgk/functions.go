package chgk

import "strings"

func maskString(s string) string {
	var sb strings.Builder
	for _, s := range []rune(strings.ToLower(s)) {
		if s >= 'а' && s <= 'я' {
			sb.WriteRune('x')
		} else if s >= 'А' && s <= 'Я' {
			sb.WriteRune('X')
		} else {
			sb.WriteRune(rune(s))
		}
	}
	return sb.String()
}
