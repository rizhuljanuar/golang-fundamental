package main

import "fmt"

func main() {
  // membuat slice langsung
  var iniSlice1 = []int{1, 2, 3, 4, 5}
  
  fmt.Println(len(iniSlice1)) // panjang length
  fmt.Println(cap(iniSlice1)) // cek kapasitas

  iniSlice2 := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}

  fmt.Println(iniSlice1)
  fmt.Println(iniSlice2)

  // membuat slice dari array
  iniArray := [8]int{10, 11, 12, 13, 14, 15, 16, 17}
  // iniSlice3 := iniArray[3:5]
  iniSlice3 := iniArray[4:]

  fmt.Printf("iniSlice3 = %v\n", iniSlice3)
  fmt.Printf("length = %d\n", len(iniSlice3))
  fmt.Printf("capacity = %d\n", cap(iniSlice3))

  // membuat slice dengan fungsi make()
  iniSlice4 := make([]int, 5, 5) // parameter 2 dan 3 itu length dan capacity
  fmt.Printf("iniSlice4 = %v\n", iniSlice4)
  fmt.Printf("length = %d\n", len(iniSlice4))
  fmt.Printf("capacity = %d\n", cap(iniSlice4))
}
