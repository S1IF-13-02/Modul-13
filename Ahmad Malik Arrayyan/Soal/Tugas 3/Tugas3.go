package main

import "fmt"

func main() {
    var target int
    fmt.Scan(&target)

    total := 0
    donatur := 0

    for total < target {
        var donasi int
        fmt.Scan(&donasi)

        donatur++
        total += donasi

        fmt.Printf(
            "Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            donatur, donasi, total,
        )
    }

    fmt.Printf(
        "Target tercapai! Total donasi: %d dari %d donatur.\n",
        total, donatur,
    )
}