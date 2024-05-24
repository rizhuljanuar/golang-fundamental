package main

import "fmt"

func main() {

  // fungsi Print() dia langsung default 1 baris tanpa spasi atau enter
  var firstName, lastName string = "rizhul", "januar"
  
  // bisa otomatis langsung ada spasi selain string
  var a, b = 100, 200

  fmt.Print(firstName, "\n")
  fmt.Print(lastName, "\n")

  fmt.Print(a, b)

  // fungsi Println() dia langsung default dengan spasi dan enter
  
  var nickName, myAddress = "rizhul", "Tangerang"

  fmt.Println(nickName, myAddress)

  var x = "Halo"
  var y = 28

  var angka = 15.5
  var txt = "hello world"

  fmt.Printf("value x adalah : %v dan tipe nya x adalah %T\n", x, x)
  fmt.Printf("value y adalah : %v dan tipe nya y adalah %T\n", y, y)

  fmt.Printf("%v%%\n", angka) // 15.5%
  fmt.Printf("%#v\n", txt) // "hello world", ada tanya kutip nya karena go syntax
}
