package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	target := math.Ceil(x) 
	current := x

	for {
		current = current + 0.1

		current = math.Round(current*10) / 10

		fmt.Println(current)

		if current >= target {
			break 
		}
	}
}
