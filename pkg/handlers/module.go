package handlers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"v2/pkg/convert"
)

func (c *Controller) ModuleHandler() error {
	if len(os.Args) < 3 {
		c.Print(ErrNotHaveProjectName)
		return errors.New(ErrNotHaveProjectName)
	}

	name := convert.ToModuleDir(os.Args[2])

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

	c.Print(fmt.Sprintf("Module %s is created successfully", name))
	return nil
}
