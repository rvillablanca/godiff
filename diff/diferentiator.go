package diff

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func CompareFiles(file1, file2 string) (bool, error) {
	size := 1024 * 4
	buffer1, buffer2 := make([]byte, size), make([]byte, size)

	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	for {
		// Si hay error, se retorna
		n1, err := f1.Read(buffer1)
		if err != nil && err != io.EOF {
			return false, err
		}
		n2, err2 := f2.Read(buffer2)
		if err2 != nil && err2 != io.EOF {
			return false, err2
		}

		// Se llegó al final de ambos archivos.
		if err == io.EOF && err == err2 && n1 == n2 {
			return true, nil
		}

		// Se leyeron distintas cantidades, se llegó al final de uno.
		if n1 != n2 {
			return false, nil
		}

		result := bytes.Equal(buffer1, buffer2)
		if !result {
			return false, nil
		}
	}
}

func FindFilesIn(list []string, dirname string) []string {
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			p, err := filepath.Rel(dirname, path)
			if err == nil {
				list = append(list, p)
			}
			return err
		}
		return nil
	}

	filepath.Walk(dirname, walkFunc)
	return list
}
