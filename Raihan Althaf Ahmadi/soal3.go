package main

import "fmt"
func main() {
    var target, donasi, total, jumlah int

    fmt.Print("Masukkan target donasi: ")
    fmt.Scan(&target)
    for {
        fmt.Print("Masukkan donasi: ")
        fmt.Scan(&donasi)

        jumlah++
        total += donasi    

        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            jumlah, donasi, total)

        if total >= target {
            break
        }
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n",
        total, jumlah)
}
