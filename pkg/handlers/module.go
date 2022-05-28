package handlers

import (
	"os"
	"path/filepath"
	"v2/pkg/converters"
)

func (c *Controller) ModuleHandler() error {
	if len(os.Args) < 3 {
		c.Print("enter a name for the new module")
		return ErrNotHaveProjectName
	}

	name := converters.Convert("", os.Args[2], "")

	if err := os.Mkdir(filepath.Join("src", "modules", name), 0777); err != nil {
		return err
	}

	if err := os.Chdir(filepath.Join("src", "modules", name)); err != nil {
		return err
	}

	os.Mkdir("actions", 0777)
	os.Mkdir("components", 0777)
	os.Mkdir("containers", 0777)
	os.Mkdir("enums", 0777)
	os.Mkdir("interfaces", 0777)
	os.Mkdir("styles", 0777)
	os.Mkdir("templates", 0777)

	return nil
}
