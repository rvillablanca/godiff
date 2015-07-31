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

	isDir1, err := diff.CheckDirectory(dir1)
	if err != nil {
		return
	}
	isDir2, err := diff.CheckDirectory(dir2)
	if err != nil {
		return
	}
	if !isDir1 || !isDir2 {
		fmt.Println("Todos los argumentos deben ser directorios")
		return
	}

	result, err := diff.CompareFiles(dir1, dir2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Resultado", result)
}

func checkArguments(args []string) bool {
	fmt.Println("Arg: ", args, len(args))
	return len(args) == 4
}

func printUsages() {
	fmt.Println("Uso: ")
	fmt.Println(os.Args[0], "<old-sources> <new-sources> <destination>")
}
