package main

import "fmt"

func main() {
    var target int
    var total int
    var donasi int
    var donatur int

    fmt.Print("Masukkan target donasi: ")
    fmt.Scan(&target)

    total = 0
    donatur = 0

    for {
        donatur = donatur + 1

        fmt.Scan(&donasi)
        total = total + donasi

        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)

        if total >= target {
            break
        }
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur)
}
	