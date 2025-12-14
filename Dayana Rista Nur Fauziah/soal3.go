package main

import "fmt"

func main() {
	var target int
	var donasi int
	total := 0
	jumlahDonatur := 0

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for total < target {
		fmt.Scan(&donasi)
		jumlahDonatur++
		total = total + donasi
		fmt.Println("Donatur", jumlahDonatur, "menyumbang", donasi, "total terkumpul", total)
	}

	fmt.Println("Target tercapai total donasi", total, "dari", jumlahDonatur, "donatur")
}
