package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"v2/pkg/fm"
)

func (c *Controller) NewHandler() error {
	if len(os.Args) < 3 {
		c.Print("enter a name for the new project")
		return ErrNotHaveProjectName
	}

	// TODO: validator/ name of project
	name := os.Args[2]

	if err := os.Mkdir(name, 0777); err != nil {
		return err
	}

	err := fm.CopyDir(filepath.Join(c.config.REACT_CLI, "template", "empty-project"), name)
	if err != nil {
		return err
	}
	c.Print(fmt.Sprintf("Project %s is created successfully", name))

	return nil
}
