package main

import "fmt"

func main() {
	var target, donasi, total, donatur int

	fmt.Print("Masukkan Target Donasi = ")
	fmt.Scan(&target)
	

	for {
		fmt.Print("Masukkan Donasi = ")
		fmt.Scan(&donasi)

		donatur++
		total += donasi

		fmt.Print("Donatur ",donatur,": Menyumbang ", donasi,". Total terkumpul: ", total, "\n")


	if total >= target {
		fmt.Println("Target tercapai! Total donasi: ", total, "dari", donatur, "donatur")
		break
	}
	}
}
