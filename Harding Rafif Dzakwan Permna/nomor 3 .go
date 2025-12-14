package main

import "fmt"

func main() {
    var target, donasi, total, jumlahDonatur int   
    
    fmt.Scan(&target)

    for total < target {
        fmt.Scan(&donasi)
        jumlahDonatur = jumlahDonatur + 1
        total = total + donasi
        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            jumlahDonatur, donasi, total)
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n",
        total, jumlahDonatur)
}
