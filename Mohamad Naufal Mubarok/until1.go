package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scan(&n)

    count := 0
    temp := n

    for {
        count++
        temp = temp / 10

        if temp == 0 {
            break
        }
    }

    fmt.Println("Jumlah digit:", count)
}