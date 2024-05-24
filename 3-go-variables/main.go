package main

import "fmt"

var x = 'ets' // bisa di eksekusi diluar fungsi

// x := 'est' // error, tidak bisa di eksekusi diluar fungsi

func main() {
  var myName string = "zhul"

  // bisa dilakukan terpisah seperti ini
  // var myName string
  // myName = "zhul"
  
  var nickName = "januar"
  myAge := 28 // tidak boleh terpisah, harus satu baris

  // deklarasi tanpa nilai awal
  var a string // ''
  var b int // 0
  var c bool // false

  // deklarasi dengan multiple variable
  var d, e, f, g int = 1, 2, 3, 4

  // deklarasi multiple variable dalam blok
  var (
    firstName string = 'rizhul'
    lastName         = 'januar'
    height    int    = 163
  )

  fmt.Println(myName, nickName, myAge)
}
