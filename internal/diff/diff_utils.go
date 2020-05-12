package diff

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rvillablanca/godiff/internal/utils"
)

var SkipKnownFolders = []string{"nbproject", ".git", ".svn", ".ant-targets-build.xml"}

// CompareFiles compara entre 2 archivos e indica si son iguales
func CompareFiles(file1, file2 string) (bool, error) {
	size := 1024 * 4
	buffer1, buffer2 := make([]byte, size), make([]byte, size)

	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	defer utils.CloseQuietly(f1)

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer utils.CloseQuietly(f2)

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

// FindFilesIn Busca archivos en el directorio dirname y retorna lista con nombres
func FindFilesIn(dirname string) []string {
	var list []string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			p, err := filepath.Rel(dirname, path)
			if err != nil {
				return err
			}
			list = append(list, p)
			return nil
		}
		return nil
	}

	_ = filepath.Walk(dirname, walkFunc)
	return list
}

// FindFilesFiltering busca archivos ignorando los subdirectorios ignored
func FindFilesFiltering(dirname string, ignored []string) []string {
	var list []string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			for _, skip := range ignored {
				if strings.Contains(path, skip) {
					return nil
				}
			}
			p, err := filepath.Rel(dirname, path)
			if err != nil {
				return err
			}
			list = append(list, p)
			return nil
		}
		return nil
	}

	_ = filepath.Walk(dirname, walkFunc)
	return list
}

func ShouldSkip(path string, ignored []string) bool {
	for _, skip := range ignored {
		if strings.Contains(path, skip) {
			return true
		}
	}
	return false
}
