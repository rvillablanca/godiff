package diff

import "testing"
import "fmt"

func TestWalk(t *testing.T) {
  fmt.Println("Ejectando test walk...")
  list := []string{}
  list = FindFilesIn(list, "/home/mechon/go")
  fmt.Println("Archivos:", list)
  if (len(list) == 0) {
    t.Error("Error en test walk")
  }
}
