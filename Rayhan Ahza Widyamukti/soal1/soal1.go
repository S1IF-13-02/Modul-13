package main

import "fmt"

func main() {
	var n int
	var hitung int = 0

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	for {
		n = n / 10
		hitung++

		if n == 0 {
			break
		}
	}

	fmt.Println("Jumlah digit:", hitung)
}
