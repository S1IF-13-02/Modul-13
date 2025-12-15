package main

import "fmt"

func main() {
	var angka float64
	fmt.Scan(&angka)

	angkaInt := int(angka * 10)
	target := ((int(angka) + 1) * 10)

	current := angkaInt
	for selesai := false; !selesai; {
		current = current + 1
		fmt.Printf("%.1f\n", float64(current)/10.0)
		selesai = current >= target
	}
}
