package convert

import (
	"strings"
	"unicode"
)

// Name - преоразует строку, возводя первую букву в верхний регситр,
// и добавляет перфикс и постфикс.
// (первую букву постфикса тоже приводит в вверхний регистр)
func Name(prefix, name, postfix string) string {
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

func ToModuleDir(name string) string {
	return Name("", name, "")
}

func ToContainerFile(name string) string {
	return Name("container", name, "")
}

func ToTemplateFile(name string) string {
	return Name("template", name, "")
}

func ToStyleFile(name string) string {
	return Name("style", name, "")
}

func ToStateName(name string) string {
	return Name("", name, "state")
}

func ToPropsName(name string) string {
	return Name("", name, "props")
}

func ToInterfaceName(name string) string {
	return Name("", name, "interface")
}

func ToClassName(name string) string {
	return Name("", name, "")
}

func ToTemplateName(name string) string {
	return Name("Template", name, "")
}

func ToStyle(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}
