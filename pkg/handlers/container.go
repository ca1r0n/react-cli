package handlers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"v2/pkg/convert"
	"v2/pkg/fill"
)

func (c *Controller) ContainerHandler() error {
	if len(os.Args) < 3 {
		c.Print(ErrNotHaveModuleName)
		return errors.New(ErrNotHaveModuleName)
	}

	if len(os.Args) < 4 {
		c.Print(ErrNotHaveContainerName)
		return errors.New(ErrNotHaveContainerName)
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
			c.Print(err.Error())
			return err
		}

		if _, err := os.Stat(pathStyle); err == nil {
			c.Print(err.Error())
			return err
		}

		if _, err := os.Stat(pathTemplate); err == nil {
			c.Print(err.Error())
			return err
		}
	}

	fileContainer, err := os.Create(pathContainer)
	if err != nil {
		c.Print(err.Error())
		return err
	}
	defer fileContainer.Close()

	fileStyle, err := os.Create(pathStyle)
	if err != nil {
		c.Print(err.Error())
		return err
	}
	defer fileStyle.Close()

	fileTemplate, err := os.Create(pathTemplate)
	if err != nil {
		c.Print(err.Error())
		return err
	}
	defer fileTemplate.Close()

	{
		if err = fill.Template(fileTemplate, container); err != nil {
			c.Print(err.Error())
			return err
		}

		if err = fill.Container(fileContainer, container); err != nil {
			c.Print(err.Error())
			return err
		}

		if err = fill.Style(fileStyle, container); err != nil {
			c.Print(err.Error())
			return err
		}
	}

	c.Print(fmt.Sprintf("Container %s is created successfully", container))
	return nil
}
