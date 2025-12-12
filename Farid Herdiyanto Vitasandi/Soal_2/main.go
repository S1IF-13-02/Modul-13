package main

import (
	"fmt"
	"math"
)

func main() {
	
	var bilangan_desimal float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&bilangan_desimal)

	sum := bilangan_desimal
	target := math.Ceil(bilangan_desimal)

	for{
		sum += 0.1
		sum = math.Round(sum*10) / 10
		fmt.Printf("%.1f\n", sum)

		if sum >= target{
			break
		}
	}
}
