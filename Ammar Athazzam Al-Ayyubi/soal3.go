package main

import "fmt"

func main() {
	var target, sumbangan int
	fmt.Print("Masukkan target donasi: ")
	total := 0
	donatur := 0
	fmt.Scan(&target)

	for {
		donatur++
		fmt.Printf("donasi ke %d: ", donatur)
		fmt.Scan(&sumbangan)
		total += sumbangan

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, sumbangan, total)

		if total >= target {
			fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur)
			break
		}
	}
}
