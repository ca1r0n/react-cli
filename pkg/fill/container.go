package fill

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
	"v2/pkg/convert"
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
		NameCss:          convert.Scss(convert.ToStyleFile(name)),
		NameState:        convert.ToStateName(name),
		NameProps:        convert.ToPropsName(name),
		NameInterface:    convert.ToInterfaceName(name),
		NameClass:        convert.ToClassName(name),
		NameTemplate:     convert.ToTemplateName(name),
		NameTemplateFile: convert.ToTemplateFile(name),
	}

	path := filepath.Join(os.Getenv("REACT_CLI"), "template", "container")

	return template.Must(template.ParseFiles(path)).Execute(w, data)
}
