package main

import "fmt"

func main() {
    var kata string
    var jumlah int

    fmt.Print("masukan kata dan jumlah pengulangan: ")
    fmt.Scanln(&kata, &jumlah)

    i := 0

    for {
        fmt.Println(kata)
        i++

        if i == jumlah {
            break
        }
    }
}
