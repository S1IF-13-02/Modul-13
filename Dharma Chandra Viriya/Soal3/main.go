package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&target)

	total := 0
	donatur := 0

	for {
		var sumbangan int
		fmt.Print("Masukkan donasi: ")
		fmt.Scanln(&sumbangan)

		donatur++
		total += sumbangan

		fmt.Printf(
			"Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			donatur, sumbangan, total,
		)

		if total >= target {
			break
		}
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, donatur,
	)
}
