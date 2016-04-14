package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"io"

	"github.com/rvillablanca/godiff/diff"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	oldDir  = kingpin.Arg("old", "Fuentes antiguos").Required().String()
	newDir  = kingpin.Arg("new", "Fuentes nuevos").Required().String()
	destDir = kingpin.Arg("dest", "Destino del parche").Required().String()
)

func main() {

	kingpin.Parse()

	err := validateDirectories()
	if err != nil {
		fmt.Println(err)
		return
	}

	oldFiles := []string{}
	newFiles := []string{}

	fmt.Println("Buscando archivos en directorios...")
	oldFiles = diff.FindFilesIn(oldFiles, *oldDir)
	newFiles = diff.FindFilesIn(newFiles, *newDir)

	toAdd := []string{}
	toRemove := []string{}

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

	toReplace := []string{}

	fmt.Println("Comparando archivos...")
	for _, v := range oldFiles {
		v1 := filepath.Join(*oldDir, v)
		v2 := filepath.Join(*newDir, v)
		equal, err := diff.CompareFiles(v1, v2)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !equal {
			toReplace = append(toReplace, v)
		}
	}

  	for _, file := range toAdd {
		srcf := filepath.Join(*newDir, file)
		dstf := filepath.Join(*destDir, file)
		dstdirs := filepath.Dir(dstf)
		err = os.MkdirAll(dstdirs, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = Copy(srcf, dstf)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, file := range toReplace {
		srcf := filepath.Join(*newDir, file)
		dstf := filepath.Join(*destDir, file)
		dstdirs := filepath.Dir(dstf)
		err = os.MkdirAll(dstdirs, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = Copy(srcf, dstf)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Eliminar los archivos:")
	for _, file := range toRemove {
		df := filepath.Join(*oldDir, file)
		fmt.Println(df)
	}

	fmt.Println("Finalizado")
}

func validateDirectories() (err error) {
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
	return nil
}

func checkDirectories(dirs ...string) (bool, error) {
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

func Copy(src, dst string) error {
  src_file, err := os.Open(src)
  if err != nil {
    return err
  }
  defer src_file.Close()

  src_file_stat, err := src_file.Stat()
  if err != nil {
    return err
  }

  if !src_file_stat.Mode().IsRegular() {
    return fmt.Errorf("%s is not a regular file", src)
  }

  dst_file, err := os.Create(dst)
  if err != nil {
    return err
  }
  defer dst_file.Close()
  _, err = io.Copy(dst_file, src_file)
	return err
}
