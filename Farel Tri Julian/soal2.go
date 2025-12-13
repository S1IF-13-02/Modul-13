package main

import (
	"fmt"
	"math"
)

func main() {
	var in float64
	fmt.Print("masukkan bilangan desimal : ")
	fmt.Scan(&in)

	batas := math.Ceil(in)

	for i := in + 0.1; i <= batas; i += 0.1 {
		if i == float64(int64(i)) {
			fmt.Printf("%.0f\n", i)
		} else {
			fmt.Printf("%.1f\n", i)
		}
	}
}
