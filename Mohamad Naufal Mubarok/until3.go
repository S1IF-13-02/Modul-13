package main

import (
    "fmt"
)

func main() {
    var target, donasi, total int
    var donorKe int

    fmt.Print("Masukkan target donasi: ")
    fmt.Scan(&target)

    for {
        donorKe++
        fmt.Printf("Donatur %d menyumbang: ", donorKe)
        fmt.Scan(&donasi)

        total += donasi
        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donorKe, donasi, total)

        
        if total >= target {
            break
        }
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donorKe)
}