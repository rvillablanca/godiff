package main

import (
	"fmt"
	"github.com/rvillablanca/godiff/diff"
	"os"
	"log"
)

func main() {
	fmt.Println("Argumentos:", os.Args)
	result, err := diff.CompareFiles(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resultado", result)
}
