package fill

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
	"v2/pkg/convert"
)

type styleTemplate struct {
	Name string
}

func Style(w io.Writer, name string) error {
	data := styleTemplate{
		Name: convert.ToStyle(name),
	}

	path := filepath.Join(os.Getenv("REACT_CLI"), "template", "containerStyle")

	return template.Must(template.ParseFiles(path)).Execute(w, data)
}
