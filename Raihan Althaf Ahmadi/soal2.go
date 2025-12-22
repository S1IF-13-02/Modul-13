package main

import (
	"fmt"
	"math"
)

func main() {
    var x float64
    fmt.Print("Masukkan bilangan desimal: ")
    fmt.Scan(&x)

    target := math.Ceil(x)

    for {
        x = math.Round((x+0.1)*10) / 10 
        fmt.Println(x)
        if x >= target {
            break
        }
    }
}
