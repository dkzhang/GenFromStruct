package utils

import (
	"strings"
	"unicode"
)

func CamelToSnake(s string) string {
	var sb strings.Builder

	var cPrev rune = 0
	var cPrevPrev rune = 0

	for i, c := range s {
		switch i {
		case 0:
			sb.WriteRune(unicode.ToLower(c))
			cPrevPrev = c
			continue
		case 1:
			cPrev = c
			continue
		default:
			if unicode.IsUpper(cPrev) && (unicode.IsLower(cPrevPrev) || unicode.IsLower(c)) {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(cPrev))
			cPrevPrev = cPrev
			cPrev = c
		}
	}
	if unicode.IsUpper(cPrev) && unicode.IsLower(cPrevPrev) {
		sb.WriteRune('_')
	}
	sb.WriteRune(unicode.ToLower(cPrev))

	return sb.String()
}

func Camel2camel(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}
