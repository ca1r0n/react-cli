package main

import (
	"os"
	"v2/pkg/handlers"
)

func main() {
	if len(os.Args) == 1 {
		handlers.NotCommandHandler()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		handlers.HelpHandler()
		os.Exit(0)
	default:
		handlers.NotFoundHandler()
		os.Exit(1)
	}
}
