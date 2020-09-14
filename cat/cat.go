package cat

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Concat reads files
func Concat(paths []string) {
	for _, file := range paths {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
