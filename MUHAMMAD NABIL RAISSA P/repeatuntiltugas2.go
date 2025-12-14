package main

import "fmt"

func main() {
	var target, donasi int
	total := 0
	jumlahDonatur := 0


	fmt.Scan(&target)

	for {
		
		fmt.Scan(&donasi)

		jumlahDonatur++
		total += donasi

		fmt.Printf(
			"Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			jumlahDonatur, donasi, total,
		)

		if total >= target {
			break
		}
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, jumlahDonatur,
	)
}
