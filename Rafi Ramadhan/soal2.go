package main

import (
    "fmt"
    "math"
)
func main() {
    var input float64
    fmt.Print("Masukkan bilangan desimal: ")
    fmt.Scan(&input)

    batasAtas := math.Ceil(input)

    for i := input + 0.1; i <= batasAtas; i += 0.1 {
        if i == float64(int64(i)) {
            fmt.Printf("%.0f\n", i)
        } else {
            fmt.Printf("%.1f\n", i)
        }
    }
}
