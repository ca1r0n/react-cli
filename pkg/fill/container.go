package fill

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
	"v2/pkg/converters"
)

type containerTemplate struct {
	NameCss          string
	NameState        string
	NameProps        string
	NameInterface    string
	NameClass        string
	NameTemplate     string
	NameTemplateFile string
}

func Container(w io.Writer, name string) error {
	data := containerTemplate{
		NameCss:          converters.Format(converters.Convert("style", name, ""), "scss"),
		NameState:        converters.Convert("", name, "state"),
		NameProps:        converters.Convert("", name, "props"),
		NameInterface:    converters.Convert("", name, "interface"),
		NameClass:        converters.Convert("", name, ""),
		NameTemplate:     converters.Convert("Template", name, ""),
		NameTemplateFile: converters.Convert("template", name, ""),
	}

	return template.Must(template.ParseFiles(filepath.Join(os.Getenv("REACT_CLI"), "template", "container"))).
		Execute(w, data)
}
