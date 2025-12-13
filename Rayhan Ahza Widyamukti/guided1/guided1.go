package main

import (
	"fmt"
)

func main() {
	var kata string
	var jumlah int
	var hitung int

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	fmt.Print("Masukkan jumlah pengulangan: ")
	fmt.Scan(&jumlah)

	hitung = 0

	for {
		fmt.Println(kata)
		hitung++

		if hitung == jumlah {
			break
		}
	}
}
