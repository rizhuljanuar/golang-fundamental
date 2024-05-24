package main

import "fmt"

const LENGTH int = 10 // const tidak bisa di pisahkan seperti variable
const WIDTH = 5

func main()  {
  fmt.Println(LENGTH)
  fmt.Println(WIDTH)

  // const tidak dapat di ubah (read-only)
  //const PI = 3.14
  //PI = 10
  
  // const bisa deklarasi multiple dalam block
  const (
    X int = 10
    Y     = 3.14
    Z     = "hello"
  )
}
