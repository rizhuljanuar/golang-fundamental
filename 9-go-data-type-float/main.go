package main

import "fmt"

func main() {
  var a float32 = 125.55
  var b float32 = 3.4e+38

  fmt.Printf("Value: %v, Tipe data: %T\n", a, a)
  fmt.Printf("Value: %v, Tipe data: %T\n", b, b)

  var c float64 = 255.55
  var d float64 = 1.7e+308

  fmt.Printf("Value: %v, Tipe data: %T\n", c, c)
  fmt.Printf("Value: %v, Tipe data: %T\n", d, d)
}
