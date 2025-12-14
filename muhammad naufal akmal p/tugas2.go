package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	atas := int(x)
	if float64(atas) != x {
		atas++
	}

	angka := x

	for {
		angka = angka + 0.1
		fmt.Printf("%.1f\n", angka)

		if angka >= float64(atas) { 
			break
		}
	}
}