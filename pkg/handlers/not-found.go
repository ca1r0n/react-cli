package handlers

import (
	"fmt"
	"os"
)

func NotFoundHandler() {
	fmt.Printf("Command %s not found\n", os.Args[1])
}
