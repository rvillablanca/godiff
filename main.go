package main

import (
	"fmt"
	"github.com/rvillablanca/godiff/diff"
	"os"
	"log"
	"errors"
)

func main() {
	args := os.Args
	err := checkArguments(args)
	if err != nil {
		printUsages()
		return
	}
	result, err := diff.CompareFiles(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resultado", result)
}

func checkArguments(args []string) error {
	if (len(args) != 4) {
		return errors.New("Número de parámetros incorrecto")
	}
	return nil
}

func printUsages() {
	fmt.Println("Uso: ")
	fmt.Println(os.Args[0], "<old-sources> <new-sources> <destination>")
}
