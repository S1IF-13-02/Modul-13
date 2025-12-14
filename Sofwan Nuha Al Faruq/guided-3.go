package main

import (
    "fmt"
)

func main() {
    var x, y int

    fmt.Print("Masukkan X: ")
    fmt.Scan(&x)
    fmt.Print("Masukkan Y: ")
    fmt.Scan(&y)

    hasil := x

    for {
        hasil = hasil - y
        fmt.Println(hasil)

        if hasil <= 0 {
            break
        }
    }

    if hasil == 0 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
