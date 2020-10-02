package copy

import (
	"fmt"
	"io"
	"os"
)

// Copy will copy args[0] to args[1]
func Copy(args []string, buffersize int) error {

	source, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(args[1])
	if err == nil {
		return fmt.Errorf("file %s already exists", args[1])
	}

	dest, err := os.Create(args[1])
	if err != nil {
		return err
	}
	defer dest.Close()

	if buffersize == 0 {
		buffersize = 10240
	}
	// create a buffer of bytes with size `buffersize`
	buf := make([]byte, buffersize)

	for {
		n, err := source.Read(buf)

		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		if _, err := dest.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
