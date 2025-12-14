package main

import (
    "fmt"
)

func main() {
    var a string
    var b int

    fmt.Print("Masukkan kata: ")
    fmt.Scan(&a)

    fmt.Print("Masukkan jumlah pengulangan: ")
    fmt.Scan(&b)

    hitung := 0

    for {
        fmt.Println(a)
        hitung++

        if hitung == b {
            break
        }
    }
}
