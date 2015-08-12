package main

import "testing"
import "fmt"

func TestIteration(t *testing.T) {
  urlList := []string{"test", "abc", "def", "ghi"}
  remove := []string{"test", "abc"}

  loop:
  for i := 0; i < len(urlList); i++ {
      url := urlList[i]
      fmt.Println("Index:", i, ",Word:", url)
      for _, rem := range remove {
        fmt.Println("Searching for:", rem)
          if url == rem {
              urlList = append(urlList[:i], urlList[i + 1:]...)
              i-- // Important: decrease index
              fmt.Println("Found, index:", i)
              continue loop
          }
      }
  }

  fmt.Println(urlList)

  if urlList[0] != "def" || urlList[1] != "ghi" {
    t.Error("Error de iteración", urlList)
  }
}
