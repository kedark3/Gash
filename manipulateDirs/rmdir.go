package manipulatedirs

import (
	"fmt"
	"os"
)

// Rmdir makes new dirs at given path
func Rmdir(args []string) {
	for _, dirname := range args {
		err := os.Remove(dirname)
		if err != nil {
			fmt.Printf("Dir %s doesn't exists or isn't empty", dirname)
		}
	}
}
