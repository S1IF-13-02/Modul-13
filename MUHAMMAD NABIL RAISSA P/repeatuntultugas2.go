package main

import "fmt"

func main() {
	var angka float64
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	batas := float64(int(angka) + 1)

	for {
		angka += 0.1

		angka = float64(int(angka*10+0.5)) / 10

		if angka > batas {
			break
		}

		fmt.Printf("%.1f\n", angka)
	}
}
