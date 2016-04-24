package main

import (
	"fmt"
	"os"

	"github.com/tecnoscimmie/libfle"
)

func splitargs() string {
	args := os.Args[1:]
	s := ""

	for i := 0; i < len(args); i++ {
		s += args[i] + " "
	}

	return s
}

func main() {
	argstring := splitargs()
	result := libfle.NewFle(argstring)

	fmt.Printf(result + "\n")
}
