package os

import "os"

// FileExists returns if a file exists or not.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
