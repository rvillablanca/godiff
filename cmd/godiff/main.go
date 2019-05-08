package main

import (
	"fmt"
	"os"

	"github.com/rvillablanca/godiff/internal/pkg/diff"
)

var (
	oldDir  = ""
	newDir  = ""
	destDir = ""
)

func main() {

	if len(os.Args) != 4 {
		_, _ = fmt.Fprint(os.Stderr, "Número de argumentos incorrectos")
		return
	}

	oldDir = os.Args[1]
	newDir = os.Args[2]
	destDir = os.Args[3]

	diff.Patch(oldDir, newDir, destDir)
}
