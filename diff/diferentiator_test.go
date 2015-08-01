package diff

import "testing"
import "fmt"

func TestWalk(t *testing.T) {
  fmt.Println("Ejecutando test walk...")
  list := []string{}
  list = FindFilesIn(list, "/home/mechon/go")
  if (len(list) == 0) {
    t.Error("Error en test walk")
  }
}
