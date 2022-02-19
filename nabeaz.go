package main
import (
  "fmt"
  "strings"
  "strconv"
)

func main() {
  for i := 1; i < 41; i++ {
      // iの値が1から40まで変化していく
      if i % 3 == 0 || strings.Contains(strconv.Itoa(i), "3") {
        fmt.Printf("アホ")
      }
      if i % 5 == 0 {
        fmt.Printf("ナルシスト")
      }
      fmt.Printf("%d\n", i)
  }
}
