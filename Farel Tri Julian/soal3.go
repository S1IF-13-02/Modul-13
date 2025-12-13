package main

import "fmt"

func main() {
	var target, total, donatur int
	fmt.Print("masukkan target donasi: ")
	fmt.Scan(&target)

	total = 0
	donatur = 0

	for {
		var donasi int
		donatur++

		fmt.Printf("Donatur %d menyumbang: ", donatur)
		fmt.Scan(&donasi)

		total += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)

		if total >= target {
			break
		}
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur)
}
