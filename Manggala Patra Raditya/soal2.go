package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&n)
	fmt.Println("Keluaran:")

	batas := math.Ceil(n)
	x := n

	for x < batas {
		x = x + 0.1
		fmt.Printf("%.1f\n", x)
	}
}