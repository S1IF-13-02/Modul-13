package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)
	batas := math.Ceil(x)
	nilai := x
	for {
		nilai += 0.1
		fmt.Printf("%.1f\n", nilai)

		if nilai >= batas {
			break
		}
	}
}
