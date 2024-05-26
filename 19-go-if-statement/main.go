package main

import "fmt"

func main() {

  // prohram if statement
  number := 0

  if number == 10 {
    fmt.Println("number adalah 10")
  } 

  // program if else statement
  angka := 5

  if angka%2 == 0 {
    fmt.Println("angka genap")
  } else {
    fmt.Println("angka ganjil")
  }

  // program else if statement
  nilai := 85
  absen := 90

  if nilai > 90 && absen > 95 {
    fmt.Println("Sangat Baik")
  } else if nilai > 80 && absen > 85 {
    fmt.Println("Baik")
  } else {
    fmt.Println("Cukup")
  }

  // program nested if statemen (if bersarang)
  num := 100

  if num >= 80 {
    fmt.Println("angka lebih dari 80")

    if num > 90 {
      fmt.Println("angka lebih dari 90")
    }
  } else {
    fmt.Println("angka kurang dari 80")
  }
}
