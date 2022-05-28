package converters

// ToContainerName - Преобразует строку в имя для контейнера(без разрешения)
func ToContainerName(name string) string {
	return Convert("container", name, "")
}

// ToStyleName - Преобразует строку в имя для стилей(без разрешения)
func ToStyleName(name string) string {
	return Convert("style", name, "")
}

// ToComponentName - Преобразует строку в имя для комонента(без разрешения)
func ToComponentName(name string) string {
	return Convert("", name, "")
}

// ToComponentStyleName - Преобразует строку в имя для стилей комонента(без разрешения)
func ToComponentStyleName(name string) string {
	return Convert("", name, "style")
}

// ToModuleName - Преобразует строку в имя для модуля(без разрешения)
func ToModuleName(name string) string {
	return Convert("", name, "")
}

func ToStateName(name string) string {
	return Convert("", name, "")
}
