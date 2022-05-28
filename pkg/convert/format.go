package converters

func Format(name, format string) string {
	return name + "." + format
}

func Tsx(name string) string {
	return Format(name, tsx)
}

func Ts(name string) string {
	return Format(name, ts)
}

func Scss(name string) string {
	return Format(name, scss)
}

func Css(name string) string {
	return Format(name, css)
}
