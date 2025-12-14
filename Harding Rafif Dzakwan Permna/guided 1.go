package main

import "fmt"

func main() {
    var kata string
    var n int

    fmt.Scan(&kata, &n)

    hitung := 0 

    for {
        fmt.Println(kata) 
		hitung++

        if hitung == n {
            break
        }
    }
}
