package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	var nilai float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	nilai = input
	batas := math.Ceil(input)

	for {
		nilai = nilai + 0.1
		fmt.Printf("%.1f\n", nilai)
		if nilai >= batas {
			break
		}
	}
}
