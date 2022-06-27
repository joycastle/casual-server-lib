package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {
	if s, err := os.Stat(path); err != nil {
		return false
	} else {
		return s.IsDir()
	}
}

//read all filenames from path with recursion
func ReadDirFiles(path string) []string {
	var filenames []string
	if !IsDir(path) {
		filenames = append(filenames, filepath.Clean(path))
		return filenames
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return filenames
	}
	for _, file := range files {
		fileName := filepath.Clean(filepath.Join(path, file.Name()))
		if IsDir(fileName) {
			filenames = append(filenames, ReadDirFiles(fileName)...)
		} else {
			filenames = append(filenames, fileName)
		}
	}

	return filenames
}
