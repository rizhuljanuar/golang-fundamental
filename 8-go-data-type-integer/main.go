package main

import "fmt"

func main() {
  // kode signed integer bisa positif dan minus
  var a int = 600
  var b int = -4600

  fmt.Printf("Tipe data: %T, value: %v\n", a, a)
  fmt.Printf("Tipe data: %T, value: %v\n", b, b)

  // kode unsigned integer hanya bisa positif
  var c uint = 600
  var d uint = 4600

  fmt.Printf("Tipe data: %T, value: %v\n", c, c)
  fmt.Printf("Tipe data: %T, value: %v\n", d, d)
}
