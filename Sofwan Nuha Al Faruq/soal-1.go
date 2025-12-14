package main

import "fmt"

func main() {
	var angka int
	var jumlahDigit int = 0

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	for {
		angka = angka / 10
		jumlahDigit++

		if angka == 0 {
			break
		}
	}

	fmt.Println("Jumlah digit:", jumlahDigit)
}
