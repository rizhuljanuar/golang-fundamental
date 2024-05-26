package main

import "fmt"

func main() {
  // contoh program single switch statement
  warnaFavorit := "biru"

  switch warnaFavorit {
    case "biru":
      fmt.Println("Warna favoritmu adalah biru")
    case "merah":
      fmt.Println("Warna favoritmu adalah merah")
    case "kuning":
      fmt.Println("Warna favoritmu adalah kuning")
    default:
      fmt.Println("Warna favoritmu adalah bukan biru, merah, dan kuning")
  }

  // contoh program multi case switch statement
  hari := "jumat"

  switch hari {
    case "senin", "selasa", "rabu", "kamis", "jumat":
      fmt.Println("Weekday")
    case "sabtu", "minggu":
      fmt.Println("Weekend")
    default:
      fmt.Println("Invalid day")
  }
}
