package pwd

import (
	"os"
)

// Pwd function returns current working directory
func Pwd() string {
	pwd, _ := os.Getwd()
	return pwd
}
