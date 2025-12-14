package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	limit := math.Ceil(x)
	value := x

	for {
		value += 0.1

		value = math.Round(value*10) / 10

		fmt.Println(value)

		if value >= limit {
			break
		}
	}
}
