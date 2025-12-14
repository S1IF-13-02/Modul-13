package main

import "fmt"

func main() {
	var targetDonasi int
	var totalTerkumpul int
	var jumlahSumbangan int
	var jumlahDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&targetDonasi)

	totalTerkumpul = 0
	jumlahDonatur = 0

	for {
		jumlahDonatur = jumlahDonatur + 1

		fmt.Scan(&jumlahSumbangan)
		totalTerkumpul = totalTerkumpul + jumlahSumbangan

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, jumlahSumbangan, totalTerkumpul)

		if totalTerkumpul >= targetDonasi {
			break
		}
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalTerkumpul, jumlahDonatur)
}
