package diff

import "testing"
import "fmt"
import "path"

func TestWalk(t *testing.T) {
  fmt.Println("Ejecutando test walk...")
  list := []string{}
  list = FindFilesIn(list, "/home/mechon/go")
  if (len(list) == 0) {
    t.Error("Error en test walk")
  }
}

func TestAbs(t *testing.T) {
  fmt.Println("Ejecutanto test de Abs...")
  falseExpected := path.IsAbs("./some/path")
  trueExpected := path.IsAbs("/dev/nul")
  if falseExpected == true || trueExpected == false {
    t.Error("Fallo en test de Abs.")
  }
}
