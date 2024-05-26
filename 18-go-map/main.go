package main

import "fmt"

func main() {
  // membuat map dengan  keyword var
  var person = map[string]string{
    "name": "rizhul januar",
    "job": "programmer",
  }

  fmt.Println(person)
  fmt.Println(person["name"])
  fmt.Println(person["job"])

  fmt.Println(len(person)) // mendapatkan panjang map

  // membuat map keyword tanda :=
  stock := map[string]int{
    "book": 10,
    "handphone": 20,
  }

  fmt.Println(stock)
  fmt.Println(stock["book"])
  fmt.Println(stock["handphone"])

  fmt.Println(len(stock)) // mendapatkan panjang map

  // membuat map dengan function make()
  var laptop = make(map[string]string)
  laptop["merk"] = "Apple"
  laptop["model"] = "Macbook Pro"

  fmt.Println(laptop)

  // menambah dan mengubah elemen map
  laptop["tahun"] = "2019"
  laptop["model"] = "Macbook Air"

  // mengahapus elemen map
  delete(laptop, "tahun")

  computer := make(map[string]int)
  computer["stock"] = 100
  computer["harga"] = 24000000

  fmt.Println(laptop)
  fmt.Println(laptop["merk"])
  fmt.Println(laptop["model"])

  fmt.Println(computer)
  fmt.Println(computer["stock"])
  fmt.Println(computer["harga"])

  // membuat map kosong
  var mapKosong1 = make(map[string]string) // nilai default nya tidak nil
  var mapKosong2 = map[string]string{} // nilai default nya nil
  
  fmt.Println(mapKosong1) // output: map[]
  fmt.Println(mapKosong2) // output: map[]

  // map bersifat pass by reference, variable utama akan berdampak ke yg lain
  var warung = map[string]string{
    "name": "Warung ucok",
    "address": "Jakarta",
  }

  warung2 := warung // terjadi disini, kita mereperesentasikan kedua variable tersebut, jika salah satu data berubah yg lain juga akan berubah

  fmt.Println("warung :", warung)
  fmt.Println("warung2 :", warung2)

  warung2["address"] = "Bekasi"

  fmt.Println("warung :", warung)
  fmt.Println("warung2 :", warung2)
}
