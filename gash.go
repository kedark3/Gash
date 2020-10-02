package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kedark3/gash/cat"
	"github.com/kedark3/gash/cd"
	"github.com/kedark3/gash/copy"
	"github.com/kedark3/gash/ls"
	manipulatedirs "github.com/kedark3/gash/manipulateDirs"
	"github.com/kedark3/gash/move"
	"github.com/kedark3/gash/pwd"
	"github.com/kedark3/gash/rm"
	"github.com/kedark3/gash/touch"
)

func main() {
	fmt.Println("Received : ", os.Args)
	switch command := os.Args[0]; {
	case strings.Contains(command, "pwd"):
		fmt.Println(pwd.Pwd())
	case strings.Contains(command, "cd"):
		cd.Cd(os.Args[1])
	case strings.Contains(command, "ls"):
		ls.List(os.Args[1:])
	case strings.Contains(command, "cat"):
		cat.Concat(os.Args[1:])
	case strings.Contains(command, "cp"):
		copy.Copy(os.Args[1:], 1024)
	case strings.Contains(command, "mv"):
		move.Move(os.Args[1:])
	case strings.Contains(command, "mkdir"):
		manipulatedirs.Mkdir(os.Args[1:])
	case strings.Contains(command, "rmdir"):
		manipulatedirs.Rmdir(os.Args[1:])
	case strings.Contains(command, "touch"):
		touch.Touch(os.Args[1:])
	case strings.Contains(command, "rm"):
		rm.Rm(os.Args[1:])
	}
}
