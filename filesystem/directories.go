package filesystem

import (
	"errors"
	"io/ioutil"
	"os"
)

// RemoveDirContents deletes the contents of a directory including
// all files and directories.
func RemoveDirContents(path string) error {
	src, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !src.IsDir() {
		return errors.New("Path is not a directory")
	}

	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range fileList {
		if file.Name()[0] == '.' {
			continue
		}

		os.Remove(path + "/" + file.Name())
	}

	return nil
}
