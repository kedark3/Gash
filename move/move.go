package move

import (
	"os"

	"github.com/kedark3/gash/copy"
)

// Move moves file from one location to another
func Move(args []string) error {
	copy.Copy(args, 1024)
	err := os.Remove(args[0])
	return err
}
