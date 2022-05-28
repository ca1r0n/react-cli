package converters

import (
	"strings"
	"unicode"
)

// Convert - преоразует строку, возводя первую букву в верхний регситр,
// и добавляет перфикс и постфикс.
// (первую букву постфикса тоже приводит в вверхний регистр)
func Convert(prefix, name, postfix string) string {
	if len(name) == 1 {
		name = string(unicode.ToUpper(rune(name[0])))
	} else {
		name = string(unicode.ToUpper(rune(name[0]))) + name[1:]
	}

	if len(postfix) != 0 {
		if len(postfix) == 1 {
			postfix = string(unicode.ToUpper(rune(postfix[0])))
		} else {
			postfix = string(unicode.ToUpper(rune(postfix[0]))) + postfix[1:]
		}
	}

	return prefix + name + postfix
}

func ToStyleName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}
