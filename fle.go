package main

import "os"
import "fmt"
import "strings"
import "math/rand"
import "github.com/fatih/color"

func getfles() []func(string) string {
	fles := []func(string) string {
  		func(d string) string {return color.New(color.FgYellow).SprintFunc()(d)},
  		func(d string) string {return color.New(color.FgBlue).SprintFunc()(d)},
  		func(d string) string {return color.New(color.FgRed).SprintFunc()(d)},
  		func(d string) string {return color.New(color.FgYellow).SprintFunc()(d)},
  		func(d string) string {return color.New(color.FgMagenta).SprintFunc()(d)},
  		func(d string) string {return color.New(color.FgCyan).SprintFunc()(d)},
  		func(d string) string {return color.New(color.FgWhite).SprintFunc()(d)},
  	}

  return fles
}

func splitargs() string {
  args := os.Args[1:]
  s := ""

  for i := 0; i < len(args); i++ {
    s += args[i]
  }

  return s
}

func main() {
  flist := getfles()

	argstring := splitargs()
	s := strings.Split(argstring, "")
	result := ""

	for i := 0; i < len(s); i++ {
		result += flist[rand.Intn(len(flist))](s[i])
	}

	fmt.Printf(result + "\n")
}
