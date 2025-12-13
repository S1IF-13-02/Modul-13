package main

import "fmt"


func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	total := 0
	donatur := 1

	for {
		var donasi int
		fmt.Scan(&donasi)

		total += donasi
		fmt.Printf("Donatur %d : Menyumbang %d. Total terkumpul: %d\n",
			donatur, donasi, total)

		donatur++

		if total >= target {
			break
		}
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, donatur-1)
}
