package manipulatedirs

import (
	"fmt"
	"os"
)

// Mkdir makes new dirs at given path
func Mkdir(args []string) {
	for _, dirname := range args {
		err := os.MkdirAll(dirname, 0755)
		if err != nil {
			fmt.Printf("Dir %s already exists", dirname)
		}
	}
}
