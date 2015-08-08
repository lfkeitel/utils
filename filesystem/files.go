package filesystem

import (
	"os"
	"path/filepath"
	"time"
)

func Touch(filename string) error {
	os.MkdirAll(filepath.Dir(filename), 0775)
	os.Chtimes(filename, time.Now(), time.Now())
	file, err := os.OpenFile(filename, os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
