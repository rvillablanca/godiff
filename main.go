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

    toAdd := []string{}
    toRemove := []string{}

	oldIteration:
    for _, oldFile := range oldFiles {
		for _, newFile := range newFiles {
			if newFile == oldFile {
				continue oldIteration
			}
		}
		toRemove = append(toRemove, oldFile)
	}
	
	newIteration:
	for _, newFile := range newFiles {
		for _, oldFile := range oldFiles {
			if newFile == oldFile {
				continue newIteration
			}
		}
		toAdd = append(toAdd, newFile)
	}
	
	fmt.Printf("Eliminar: %v\n", toRemove)
	fmt.Printf("Agregar: %v\n", toAdd)
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
	dir3, err := filepath.Abs(*destDir)
	if err != nil {
		err = errors.New("No fue posible verificar el directorio" + *destDir)
		return
	}
	
	valid, err := checkDirectories(dir1, dir2, dir3)
	if err != nil {
		return
	} else if !valid {
		err = errors.New("Todos los argumentos deben ser directorios")
		return
	}
	return diffconf{dir1, dir2, *destDir}, nil
}

func checkDirectories(dirs... string) (bool, error) {
	for _, v := range dirs {
		exist, err := checkDirectory(v)
		if err != nil || !exist {
			return false, fmt.Errorf("Directorio %v no existe", v)
		}
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
