package main

import "fmt"

func main() {
	var bilangan float64
	var nilaiAkhir float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&bilangan)

	nilaiAkhir = float64(int(bilangan) + 1)

	for {
		bilangan = bilangan + 0.1
		fmt.Printf("%.1f\n", bilangan)

		if bilangan >= nilaiAkhir {
			break
		}
	}
}
