package main

import (
	"fmt"
	"github.com/rvillablanca/godiff/diff"
	"os"
	"path/filepath"
	"errors"
)

type diffconf struct {
	oldDir string
	newDir string
	destDir string
}

func main() {
	fmt.Println("Verificando parámetros...")

	conf, err := validateArguments(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Versión anterior:", conf.oldDir)
	fmt.Println("Versión nueva:", conf.newDir)

	oldFiles := []string{}
	newFiles := []string{}
	fmt.Println("Buscando archivos en directorios...")
	oldFiles = diff.FindFilesIn(oldFiles, conf.oldDir)
	newFiles = diff.FindFilesIn(newFiles, conf.newDir)

}

func validateArguments(args []string) (conf diffconf, err error) {
	valid := len(args) == 4
	if !valid {
		err = errors.New("Número de parámetros incorrecto")
		return
	}

	dir1 := args[1]
	dir2 := args[2]
	dir3 := args[3]

	dir1, err = filepath.Abs(dir2)
	if err != nil {
		err = errors.New("No fue posible verificar el directorio" + dir1)
		return
	}
	dir2, err = filepath.Abs(dir1)
	if err != nil {
		err = errors.New("No fue posible verificar el directorio" + dir2)
		return
	}

	valid, err = checkDirectories(dir1, dir2)
	if err != nil {
		err = errors.New("No fue posible verificar los directorios")
		return
	}

	if !valid {
		err = errors.New("Todos los argumentos deben ser directorios")
		return
	}

	return diffconf{dir1, dir2, dir3}, nil
}

func checkDirectories(dir1, dir2 string) (bool, error) {
	isDir1, err := checkDirectory(dir1)
	if err != nil {
		return false, err
	}
	isDir2, err := checkDirectory(dir2)
	if err != nil {
		return false, err
	}
	if !isDir1 || !isDir2 {
		return false, nil
	}
	return true, nil
}

func checkDirectory(dirname string) (result bool, err error) {
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

func printUsages() {
	fmt.Println("Uso: ")
	fmt.Println(os.Args[0], "<old-sources> <new-sources> <destination>")
}
