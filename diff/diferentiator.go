package diff

import (
	"os"
	"bytes"
	"path/filepath"
)

func CheckDirectory(dirname string) (result bool, err error) {
	f, err := os.Open(dirname)
	if err != nil {
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return
	}

	result = fi.Mode().IsDir()
	return
}

func CompareFiles(file1, file2 string) (result bool, err error) {
	size := 1024 * 8
	buffer1, buffer2 := make([]byte, size), make([]byte, size)

	f1, err := os.Open(file1)
	if err != nil {
		return
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return
	}
	defer f2.Close()

	n1, err := f1.Read(buffer1)
	if err != nil {
		return
	}
	n2, err := f2.Read(buffer2)
	if err != nil {
		return
	}

	if n1 != n2 {
		result = false
		return
	}

	result = bytes.Equal(buffer1, buffer2)
	return
}

func FindFilesIn(list []string, dirname string) []string  {
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			list = append(list, info.Name())
		}
		return nil
	}

	filepath.Walk(dirname, walkFunc)
	return list
}
