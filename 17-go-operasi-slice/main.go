package main

import "fmt"

func main() {
  // mengakses elemen slice
  numbers := []int{1, 2, 3}

  fmt.Println(numbers[0])
  fmt.Println(numbers[2])

  // mengubah elemen slice
  numbers2 := []int{4,5,6}
  numbers2[2] = 7

  fmt.Println(numbers2)

  // menambah elemen slice
  iniSlice := []int{10, 20, 30, 40, 50}
  fmt.Printf("iniSlice = %v\n", iniSlice)
  fmt.Printf("length = %d\n", len(iniSlice))
  fmt.Printf("capacity = %d\n", cap(iniSlice))

  iniSlice = append(iniSlice, 60, 70)
  fmt.Printf("iniSlice = %v\n", iniSlice)
  fmt.Printf("length = %d\n", len(iniSlice))
  fmt.Printf("capacity = %d\n", cap(iniSlice))

  // mengubah panjang slice
  iniArray := [6]int{9, 10, 11, 12, 13, 14} // array
  iniSliceArray := iniArray[1:5] // slice array
  fmt.Printf("iniSliceArray = %v\n", iniSliceArray)
  fmt.Printf("length = %d\n", len(iniSliceArray))
  fmt.Printf("capacity = %d\n", cap(iniSliceArray))

  iniSliceArray = iniArray[1:3] // mengubah length dengan re-slicing sebuah array
  fmt.Printf("iniSliceArray = %v\n", iniSliceArray)
  fmt.Printf("length = %d\n", len(iniSliceArray))
  fmt.Printf("capacity = %d\n", cap(iniSliceArray))

  iniSliceArray = append(iniSliceArray, 20, 21, 22, 23) // mengubah panjang slice dengan menambah elemen
  fmt.Printf("iniSliceArray = %v\n", iniSliceArray)
  fmt.Printf("length = %d\n", len(iniSliceArray))
  fmt.Printf("capacity = %d\n", cap(iniSliceArray))

  // copy slice
  angka := []int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println("Angka:", angka)
  fmt.Println("Length:", len(angka))
  fmt.Println("Capacity:", cap(angka))

  copyAngka := make([]int, len(angka), cap(angka))
  copy(copyAngka, angka)
  fmt.Println("Copy Angka:", copyAngka)
  fmt.Println("Length:", len(copyAngka))
  fmt.Println("Capacity:", cap(copyAngka))
}
