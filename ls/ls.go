package ls

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// List contents of given path
func List(args []string) {
	var filesOrDirs []string
	var recursive, hidden, longListing bool
	for _, s := range args {
		if s == "-R" {
			recursive = true
		}
		if s == "-a" {
			hidden = true
		}
		if s == "-al" || s == "-la" {
			// TODO: Still implement long listing
			longListing = true
			hidden = true
		}
		if !strings.HasPrefix(s, "-") {
			filesOrDirs = append(filesOrDirs, s)
		}
	}

	for _, s := range filesOrDirs {
		fi, err := os.Stat(s)
		if err != nil {
			log.Fatal(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			printDirContents(s, hidden, recursive, longListing)
		case mode.IsRegular():
			printFile(s)

		}
	}

}

func printDirContents(path string, hidden, recursive, longListing bool) {
	if recursive {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			fmt.Println(path)
			return nil
		})
		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", path, err)
			return
		}
	} else {
		dirContent, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
			return
		}
		for _, file := range dirContent {
			if strings.HasPrefix(file.Name(), ".") && !hidden {
				continue
			}
			fmt.Println(file.Name())
		}
	}

}

func printFile(path string) {
	file, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file.Name())
}
