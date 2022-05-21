package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"v2/config"
	"v2/pkg/handlers"
)

func main() {
	conf := config.LoadEnv()

	logErr, err := os.OpenFile(filepath.Join(conf.REACT_CLI, "react-cli.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	logger := log.New(logErr, "react-cli", log.Lshortfile|log.Ldate)
	controller := handlers.New(conf)

	if len(os.Args) == 1 {
		handlers.NotCommandHandler()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		if err := controller.NewHandler(); err != nil {
			fmt.Println(err)
			logger.Println(err)
		}

	case "help":
		controller.HelpHandler()
	default:
		controller.NotFoundHandler()
	}

	os.Exit(0)
}
