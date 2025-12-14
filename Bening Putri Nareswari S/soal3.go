package main

import (
	"fmt"
)

func main() {
	var target int
	var donasi int
	var totalDonasi int = 0
	var jumlahDonatur int = 0

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for {
		fmt.Printf("Masukkan donasi Donatur %d: ", jumlahDonatur+1)
		fmt.Scan(&donasi)

		totalDonasi += donasi
		jumlahDonatur++

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
		if totalDonasi >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
