package handlers

import "errors"

var (
	ErrNotHaveProjectName = errors.New("command \"new\" must contain the name of the project")
)
