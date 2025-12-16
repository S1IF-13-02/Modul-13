package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukkan bilangan pertama: ")
    fmt.Scanln(&x)
    fmt.Print("Masukkan bilangan kedua: ")
    fmt.Scanln(&y)

    if y == 0 {
        fmt.Println("Bilangan kedua tidak boleh nol.")
        return
    }

    if x%y == 0 {
        fmt.Println(x, "adalah kelipatan dari", y)
    } else {
        fmt.Println(x, "bukan kelipatan dari", y)
    }
}
