package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    hitung := 0
    for n > 0 {
        n = n / 10   
        hitung++     
    }

    fmt.Println(hitung)
}
