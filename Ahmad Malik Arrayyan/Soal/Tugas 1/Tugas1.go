package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    count := 0

    for n > 0 {
        n = n / 10
        count++
    }

    fmt.Println(count)
}
