package handlers

func (c *Controller) HelpHandler() error {
	c.Print(`Usage:
  $ react-cli [command] [flags]

Available commands:
  new       [project_name]                  	Create new project
  module/m    [module_name]                   	Create new module
  container/c [module_name] [container_name]  	Create new container in module
  help                                      	Show any commands`)
	return nil
}
