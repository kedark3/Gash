package main

import (
	"fmt"
	"os"

	"github.com/kedark3/gash/cat"
	"github.com/kedark3/gash/cd"
	"github.com/kedark3/gash/ls"
	"github.com/kedark3/gash/pwd"
)

func main() {
	switch os.Args[0] {
	case "cwd":
		fmt.Println(pwd.Pwd())
	case "chdir":
		cd.Cd(os.Args[1])
	case "list":
		ls.List(os.Args[1:])
	case "concat":
		cat.Concat(os.Args[1:])
	}
}
