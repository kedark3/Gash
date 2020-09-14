package cd

import "os"

// Cd function changes directoy to dir you want
func Cd(dirname string) {
	os.Chdir(dirname)
}
