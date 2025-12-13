package main

import "fmt"

func main() {
	var x float64
	var batas float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)

	batas = float64(int(x) + 1)

	for {
		x = x + 0.1
		fmt.Printf("%.1f\n", x)

		if x >= batas {
			break
		}
	}
}
