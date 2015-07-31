package diff

import (
	"os"
	"bytes"
)

func CompareFiles(file1, file2 string) (result bool, err error) {
	size := 1024 * 8
	var buffer1, buffer2 [size]uint8
	
	f1, err := os.Open(file1)
	if err != nil {
		return
	}
	f2, err := os.Open(file2)
	if err != nil {
		return
	}
	
	n1, err := f1.Read(buffer1)
	if err != nil {
		return
	}
	n2, err := f2.Read(buffer2)
	if err != nil {
		return
	}
	
	if n1 != n2 {
		result = false
		return
	}
	
	result = bytes.Equal(buffer1, buffer2)
	return
}
