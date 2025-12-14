package main

import "fmt"

func main() {
    var input float64
    fmt.Scan(&input)

    x := int(input * 10)
    batas := (int(input) + 1) * 10

    for x < batas {
        x += 1 

        if x == batas {
            fmt.Println(x / 10) 
        } else {
            fmt.Printf("%.1f\n", float64(x)/10)
        }
    }
}
