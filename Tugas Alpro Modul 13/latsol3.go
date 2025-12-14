package main

import "fmt"

func main() {
	var target int
	var donasi int
	total := 0
	donatur := 0

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for total < target {
		fmt.Print("Masukkan total donasi dari donatur: ")
		fmt.Scan(&donasi)

		donatur = donatur + 1
		total = total + donasi

		fmt.Println("Donatur", donatur, ": Menyumbang", donasi,
			". Total terkumpul:", total)
	}

	fmt.Println("Target tercapai! Total donasi:", total,
		"dari", donatur, "donatur.")
}
