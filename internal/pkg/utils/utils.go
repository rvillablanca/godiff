package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ValidateDirectories(oldDir, newDir, destDir string) error {
	dir1, err := filepath.Abs(oldDir)
	if err != nil {
		return errors.New("no fue posible verificar el directorio" + oldDir)
	}
	dir2, err := filepath.Abs(newDir)
	if err != nil {
		return errors.New("no fue posible verificar el directorio" + newDir)
	}
	dir3, err := filepath.Abs(destDir)
	if err != nil {
		return errors.New("no fue posible verificar el directorio" + destDir)
	}

	valid, err := checkDirectories(dir1, dir2, dir3)
	if err != nil {
		return err
	} else if !valid {
		return errors.New("todos los argumentos deben ser directorios")
	}
	return nil
}

func checkDirectories(dirs ...string) (bool, error) {
	for _, v := range dirs {
		exist, err := checkDirectory(v)
		if err != nil || !exist {
			return false, fmt.Errorf("directorio %v no existe", v)
		}
	}
	return true, nil
}

func checkDirectory(dirname string) (result bool, err error) {
	f, err := os.Open(dirname)
	if err != nil {
		return
	}
	defer CloseQuietly(f)
	fi, err := f.Stat()
	if err != nil {
		return
	}

	result = fi.Mode().IsDir()
	return
}

func Copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer CloseQuietly(srcFile)

	srcFileStat, err := srcFile.Stat()
	if err != nil {
		return err
	}

	if !srcFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer CloseQuietly(dstFile)
	_, err = io.Copy(dstFile, srcFile)
	return err
}

func CloseQuietly(c io.Closer) {
	_ = c.Close()
}
