package handlers

import (
	"errors"
	"os"
	"path/filepath"
	"v2/pkg/convert"
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

	module := convert.ToModuleDir(os.Args[2])
	container := os.Args[3]

	nameContainer := convert.Tsx(convert.ToContainerFile(container))
	nameTemplate := convert.Tsx(convert.ToTemplateFile(container))
	nameStyle := convert.Scss(convert.ToStyleFile(container))

	pathStyle := filepath.Join("src", "modules", module, "styles", nameStyle)
	pathContainer := filepath.Join("src", "modules", module, "containers", nameContainer)
	pathTemplate := filepath.Join("src", "modules", module, "templates", nameTemplate)

	{
		if _, err := os.Stat(pathContainer); err == nil {
			return errors.New("file already exist")
		}

		if _, err := os.Stat(pathStyle); err == nil {
			return errors.New("file already exist")
		}

		if _, err := os.Stat(pathTemplate); err == nil {
			return errors.New("file already exist")
		}
	}

	fileContainer, err := os.Create(pathContainer)
	if err != nil {
		return err
	}
	defer fileContainer.Close()

	fileStyle, err := os.Create(pathStyle)
	if err != nil {
		return err
	}
	defer fileStyle.Close()

	fileTemplate, err := os.Create(pathTemplate)
	if err != nil {
		return err
	}
	defer fileTemplate.Close()

	{
		if err = fill.Template(fileTemplate, container); err != nil {
			return err
		}

		if err = fill.Container(fileContainer, container); err != nil {
			return err
		}

		if err = fill.Style(fileStyle, container); err != nil {
			return err
		}
	}

	return nil
}
