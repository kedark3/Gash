package touch

import (
	"fmt"
	"os"
)

// Touch creates new files
func Touch(args []string) {
	for _, item := range args {
		emptyFile, err := os.Create(item)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("created file", emptyFile)
		emptyFile.Close()
	}
}
