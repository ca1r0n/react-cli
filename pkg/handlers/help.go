package handlers

import (
	"fmt"
)

func HelpHandler() {
	fmt.Print(`Usage:
  $ react-cli [command] [flags]

Available commands:
  project, p       Create new react project
  container, c     Create new container
  entity, e        Create new model
  component, co    Create new component
  version, v       Show version application
  help             Show any commands

Available flags:
  --help, -h       Show about commands
`)
}
