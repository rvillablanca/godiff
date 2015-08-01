package main

import (
	"fmt"
	"github.com/rvillablanca/godiff/diff"
	"os"
)

func main() {
	args := os.Args
	valid := checkArguments(args)
	if !valid {
		fmt.Println("Número de parámetros incorrecto")
		printUsages()
		return
	}
	dir1 := os.Args[1]
	dir2 := os.Args[2]

	valid, err := checkDirectories(dir1, dir2)
	if err != nil {
		fmt.Println("No fue posible verificar los directorios")
		return
	}

	if !valid {
		fmt.Println("Todos los argumentos deben ser directorios absolutos")
	}

	fmt.Println("Versión anterior:", dir1)
	fmt.Println("Versión nueva:", dir2)

	oldFiles := []string{}
	newFiles := []string{}
	fmt.Println("Buscando archivos en directorios...")
	oldFiles = diff.FindFilesIn(oldFiles, dir1)
	newFiles = diff.FindFilesIn(newFiles, dir2)

}

func checkArguments(args []string) bool {
	return len(args) == 4
}

func checkDirectories(dir1, dir2 string) (bool, error) {
	isDir1, err := diff.CheckDirectory(dir1)
	if err != nil {
		return false, err
	}
	isDir2, err := diff.CheckDirectory(dir2)
	if err != nil {
		return false, err
	}
	if !isDir1 || !isDir2 {
		return false, nil
	}
	return true, nil
}

func printUsages() {
	fmt.Println("Uso: ")
	fmt.Println(os.Args[0], "<old-sources> <new-sources> <destination>")
}
