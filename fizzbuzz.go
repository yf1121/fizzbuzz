package main
import "fmt"

func main() {
  for i := 1; i < 101; i++ {
      // iの値が1から100まで変化していく
      if i % 15 == 0 {
        println("FizzBuzz")
      } else if i % 3 == 0 {
        println("Fizz")
      } else if i % 5 == 0 {
        println("Buzz")
      } else {
        fmt.Printf("%d\n", i)
      }
  }
}
