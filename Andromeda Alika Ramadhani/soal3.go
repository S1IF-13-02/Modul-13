package main

import "fmt"

func main() {
	var target, donasi, total, jmlDonatur int
	fmt.Scan(&target)

	for {
		fmt.Scan(&donasi)
		jmlDonatur++
		total += donasi

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jmlDonatur, donasi, total)

		if total >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, jmlDonatur)
}
