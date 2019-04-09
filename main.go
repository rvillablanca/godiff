package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/rvillablanca/godiff/diff"
)

var (
	oldDir  = ""
	newDir  = ""
	destDir = ""
)

func main() {

	if len(os.Args) != 4 {
		fmt.Fprint(os.Stderr, "Número de argumentos incorrectos")
		return
	}

	oldDir = os.Args[1]
	newDir = os.Args[2]
	destDir = os.Args[3]

	err := validateDirectories()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	fmt.Println("Buscando archivos en directorios...")
	var oldFiles = diff.FindFilesIn(oldDir)
	var newFiles = diff.FindFilesIn(newDir)

	toAdd := make([]string, 0)
	toRemove := make([]string, 0)
	toReplace := make([]string, 0)

	//Archivos a eliminar
oldIteration:
	for _, oldFile := range oldFiles {
		for _, newFile := range newFiles {
			if newFile == oldFile {
				continue oldIteration
			}
		}
		toRemove = append(toRemove, oldFile)
	}

	//Archivos que hay que agregar
newIteration:
	for _, newFile := range newFiles {
		for _, oldFile := range oldFiles {
			if newFile == oldFile {
				continue newIteration
			}
		}
		toAdd = append(toAdd, newFile)
	}

	//Se quitan los archivos a eliminar de la lista de archivos a comparar.
loop:
	for i := 0; i < len(oldFiles); i++ {
		old := oldFiles[i]
		for _, rem := range toRemove {
			if old == rem {
				oldFiles = append(oldFiles[:i], oldFiles[i+1:]...)
				i--
				continue loop
			}
		}
	}

	fmt.Println("Comparando archivos...")
	for _, v := range oldFiles {
		v1 := filepath.Join(oldDir, v)
		v2 := filepath.Join(newDir, v)
		equal, compError := diff.CompareFiles(v1, v2)
		if compError != nil {
			log.Fatal("Error al comparar archivos viejos y nuevos ", err)
			return
		}
		if !equal {
			toReplace = append(toReplace, v)
		}
	}

	for _, file := range toAdd {
		srcf := filepath.Join(newDir, file)
		dstf := filepath.Join(destDir, file)
		dstdirs := filepath.Dir(dstf)
		err = os.MkdirAll(dstdirs, os.ModePerm)
		if err != nil {
			log.Fatal("Error al crear directorios para copiar los archivos nuevos ", err)
			return
		}
		err = copy(srcf, dstf)
		if err != nil {
			log.Fatal("Error al copiar archivo ", srcf, err)
			return
		}
	}

	for _, file := range toReplace {
		srcf := filepath.Join(newDir, file)
		dstf := filepath.Join(destDir, file)
		dstdirs := filepath.Dir(dstf)
		err = os.MkdirAll(dstdirs, os.ModePerm)
		if err != nil {
			log.Fatal("Error al crear directorios para copiar los archivos a reemplazar ", err)
			return
		}
		err = copy(srcf, dstf)
		if err != nil {
			log.Fatal("Error al copiar archivo ", srcf, err)
			return
		}
	}
	//Se crea archivo con lista de archivos a eliminar sólo si aplica
	if len(toRemove) > 0 {
		fileName := filepath.Join(destDir, "to_delete.txt")
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Error al crear lista de archivos a eliminar ", err)
		}
		defer f.Close()
		for _, file := range toRemove {
			f.WriteString(file + "\n")
		}
		fmt.Println("Se deben eliminar los archivos descritos en to_delete.txt")
	}

	fmt.Println("Finalizado")
}

func validateDirectories() (err error) {
	dir1, err := filepath.Abs(oldDir)
	if err != nil {
		err = errors.New("no fue posible verificar el directorio" + oldDir)
		return
	}
	dir2, err := filepath.Abs(newDir)
	if err != nil {
		err = errors.New("no fue posible verificar el directorio" + newDir)
		return
	}
	dir3, err := filepath.Abs(destDir)
	if err != nil {
		err = errors.New("no fue posible verificar el directorio" + destDir)
		return
	}

	valid, err := checkDirectories(dir1, dir2, dir3)
	if err != nil {
		return
	} else if !valid {
		err = errors.New("todos los argumentos deben ser directorios")
		return
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
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return
	}

	result = fi.Mode().IsDir()
	return
}

func copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

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
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	return err
}
