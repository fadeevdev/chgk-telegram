package chgk

import "strings"

func maskString(s string) string {
	var sb strings.Builder
	for _, s := range []rune(strings.ToLower(s)) {
		if (s >= 'а' && s <= 'я') || (s >= 'a' && s <= 'z') {
			sb.WriteRune('x')
		} else if (s >= 'А' && s <= 'Я') || (s >= 'A' && s <= 'Z') {
			sb.WriteRune('X')
		} else {
			sb.WriteRune(s)
		}
	}
	return sb.String()
}
