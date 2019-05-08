package diff

import (
	"testing"
)
import "fmt"
import "path"

func TestWalk(t *testing.T) {
	fmt.Println("executing walk test...")
	list := FindFilesIn("./")
	if len(list) == 0 {
		t.Error("error while executing walk test")
	}
}

func TestAbs(t *testing.T) {
	fmt.Println("executing test of abs ...")
	falseExpected := path.IsAbs("./some/path")
	trueExpected := path.IsAbs("/dev/nul")
	if falseExpected == true || trueExpected == false {
		t.Error("abs test failed")
	}
}
