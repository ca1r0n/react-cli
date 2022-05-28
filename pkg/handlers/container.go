package handlers

import (
	"errors"
	"os"
	"path/filepath"
	"v2/pkg/converters"
	"v2/pkg/fill"
)

func (c *Controller) ContainerHandler() error {
	if len(os.Args) < 3 {
		c.Print("enter a module name")
		return ErrNotHaveProjectName
	}

	if len(os.Args) < 4 {
		c.Print("enter a module name")
		return ErrNotHaveProjectName
	}

	nameModule := os.Args[2]
	nameContainer := os.Args[3]
	//currentPath, _ := os.Getwd()
	name := converters.ToContainerName(nameContainer) + ".tsx"

	path := filepath.Join("src", "modules", nameModule, "container", name)

	_, err := os.Stat(path)
	if err == nil {
		return errors.New("file already exist")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = fill.Container(file, nameContainer)
	if err != nil {
		return err
	}

	return nil
}
