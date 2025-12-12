package main

import "fmt"

func main() {
    var target, donasi, total, donorCount int

    fmt.Scan(&target)

    for { 
        fmt.Scan(&donasi)

        donorCount++
        total += donasi

        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            donorCount, donasi, total)

        if total >= target { 
            break
        }
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n",
        total, donorCount)
}
