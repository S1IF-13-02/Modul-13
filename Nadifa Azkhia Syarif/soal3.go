package main

import "fmt"

func main() {
	var target int
	fmt.Scan(&target)

	total := 0
	donatur := 0

	for {
		var donasi int
		fmt.Scan(&donasi)

		donatur++
		total = total + donasi

		fmt.Println("Donatur", donatur, ": Menyumbang", donasi, ". Total terkumpul:", total)

		if total >= target {
			break
		}
	}

	fmt.Println("Target tercapai! Total donasi:", total, "dari", donatur, "donatur.")
}
