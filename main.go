package main

import (
	"fmt"
	"os"
	"path/filepath"
	"errors"
	"github.com/rvillablanca/godiff/diff"
	"gopkg.in/alecthomas/kingpin.v2"
)

type diffconf struct {
	oldDir string
	newDir string
	destDir string
}

var (
	oldDir = kingpin.Arg("old", "Fuentes antiguos").Required().String()
	newDir = kingpin.Arg("new", "Fuentes nuevos").Required().String()
	destDir = kingpin.Arg("dest", "Destino del parche").Required().String()
)

func main() {
	kingpin.Parse()

	conf, err := generateAbsoluteDirectories()
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

func generateAbsoluteDirectories() (conf diffconf, err error) {
	dir1, err := filepath.Abs(*oldDir)
	if err != nil {
		err = errors.New("No fue posible verificar el directorio" + *oldDir)
		return
	}
	dir2, err := filepath.Abs(*newDir)
	if err != nil {
		err = errors.New("No fue posible verificar el directorio" + *newDir)
		return
	}
	valid, err := checkDirectories(dir1, dir2)
	if err != nil {
		err = errors.New("No fue posible verificar los directorios")
		return
	} else if !valid {
		err = errors.New("Todos los argumentos deben ser directorios")
		return
	}
	return diffconf{dir1, dir2, *destDir}, nil
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
