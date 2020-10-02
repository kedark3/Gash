package rm

import (
	"fmt"
	"os"
)

// Rm Removes files/dirs
func Rm(args []string) {
	for _, item := range args {
		err := os.RemoveAll(item)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
