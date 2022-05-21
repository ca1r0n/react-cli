package handlers

import (
	"fmt"
	"os"
)

func (c *Controller) NotFoundHandler() error {
	c.Print(fmt.Sprintf(`Command %s not found
Use "react-cli help" for more information about commands.`, os.Args[1]))
	return nil
}
