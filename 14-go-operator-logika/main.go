package main

import "fmt"

func main() {
  var a = 15

  fmt.Println(a > 10 && a < 5)
  fmt.Println(a > 10 || a < 5)
  fmt.Println(! (a > 10 || a < 5))
}
