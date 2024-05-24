package main

import "fmt"

func main() {
  // opsi 1
  var a1 = [3]int{100, 200, 300}
  a2 := [4]int{400, 500, 600, 700}

  fmt.Println(a1)
  fmt.Println(a2)

  // opsi 2
  var a3 = [...]int{10, 20, 30}
  a4 := [...]int{40, 50, 60, 70}

  fmt.Println(a3)
  fmt.Println(a4)

  // mengakses elemen array
  prices := [3]int{1000, 2000, 3000}
  // mengubah elemen array
  prices[2] = 4000

  fmt.Println(prices[0])
  fmt.Println(prices[2])

  // Initialize array
  a5 := [4]int{}        // Tidak diinisialisasi
  a6 := [4]int{1, 2}    // Diinisialisasi sebagian
  a7 := [4]int{1, 2, 3, 4} // Diinisialisasi semuanya
  
  fmt.Println(a5)
  fmt.Println(a6)
  fmt.Println(a7)

  // Mendapatkan panjang array
  a8 := [4]int{400, 500, 600, 700}
  var languanges = [3]string{"Go", "Java", "C"}

  fmt.Println(len(a8))
  fmt.Println(len(languanges))
}
