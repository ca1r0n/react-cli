package fill

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
	"v2/pkg/converters"
)

type templateTemplate struct {
	NameContainer     string
	NameContainerFile string
	NameTemplate      string
	NameInterface     string
	NameStyle         string
}

func Template(w io.Writer, name string) error {
	data := templateTemplate{
		NameContainer:     converters.Convert("", name, ""),
		NameContainerFile: converters.Convert("container", name, ""),
		NameTemplate:      converters.Convert("Template", name, ""),
		NameInterface:     converters.Convert("", name, "interface"),
		NameStyle:         converters.ToStyleName(name),
	}

	path := filepath.Join(os.Getenv("REACT_CLI"), "template", "template")

	return template.Must(template.ParseFiles(path)).Execute(w, data)
}
