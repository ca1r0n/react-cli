package fill

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
	"v2/pkg/convert"
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
		NameContainerFile: convert.ToContainerFile(name),
		NameTemplate:      convert.ToTemplateName(name),
		NameInterface:     convert.ToInterfaceName(name),
		NameStyle:         convert.ToStyle(name),
	}

	path := filepath.Join(os.Getenv("REACT_CLI"), "template", "template")

	return template.Must(template.ParseFiles(path)).Execute(w, data)
}
