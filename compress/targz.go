package compress

/**
 * The code for this package was taken from MadCrazy's question on StackOverflow
 * URL: http://stackoverflow.com/questions/13611100/how-to-write-a-directory-not-just-the-files-in-it-to-a-tar-gz-file-in-golang
 */

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func tarGzWrite(_path string, tw *tar.Writer, fi os.FileInfo) error {
	fr, err := os.Open(_path)
	if err != nil {
		return err
	}
	defer fr.Close()

	h := new(tar.Header)
	h.Name = _path
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()

	err = tw.WriteHeader(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, fr)
	if err != nil {
		return err
	}
	return nil
}

func iterDirectory(dirPath string, tw *tar.Writer) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()
	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		curPath := dirPath + "/" + fi.Name()
		if fi.IsDir() {
			iterDirectory(curPath, tw)
		} else {
			tarGzWrite(curPath, tw, fi)
		}
	}
	return nil
}

// TarGz creates a gzipped tarball named outFilePath from the directory inPath
func TarGz(outFilePath string, inPath string) error {
	// file write
	fw, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer fw.Close()

	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

	iterDirectory(inPath, tw)
	return nil
}
