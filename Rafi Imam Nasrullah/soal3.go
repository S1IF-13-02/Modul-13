package main

import "fmt"

func main() {
	var target int
	var donatur int
	var donasi int
	var total int

	fmt.Print("Target Donasi : ")
	fmt.Scan(&target)

	donatur = 0
	total = 0

	for {
		donatur++
		fmt.Printf("Donatur ke-%d Menyumbang : ", donatur)
		fmt.Scan(&donasi)
		total += donasi
		if total >= target {
			break
	}
}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur\n", total, donatur)
}