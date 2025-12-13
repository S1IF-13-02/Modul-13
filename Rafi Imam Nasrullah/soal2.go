package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	target := math.Ceil(x)
	n := x

	for {
		n += 0.1

		n = math.Round(n*10) / 10

		fmt.Printf("%.1f\n", n)

		if n >= target {
			break
		}
	}
}
