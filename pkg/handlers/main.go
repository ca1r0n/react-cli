package handlers

import (
	"fmt"
	"v2/config"
)

type Controller struct {
	config *config.Config
}

func New(config *config.Config) *Controller {
	return &Controller{
		config: config,
	}
}

func (c *Controller) Print(message string) {
	fmt.Println(message)
}
