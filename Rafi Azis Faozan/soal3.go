package main

import "fmt"

func main() {
	var donasi int
	var target int
	var donatur int
	var total int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)
	donatur = 0
	total = 0
	for {
		fmt.Print("Masukkan donasi: ")
		fmt.Scan(&donasi)
		total = total + donasi
		donatur++
		fmt.Println("Donatur", donatur, "menyumbang", donasi, "Total terkummpul", total)
		if total >= target {
			fmt.Println("Target tercapai! Total donasi:", total, "dari", donatur, "donatur")
			break
		}
	}
}
