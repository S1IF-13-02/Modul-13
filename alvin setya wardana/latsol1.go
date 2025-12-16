package main

import "fmt"

func main() {
	var n int
	var jumlahDigit int
	var selesai bool

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	jumlahDigit = 0
	for selesai = false; !selesai; {
		jumlahDigit++
		n = n / 10
		selesai = n == 0
	}

	fmt.Println(jumlahDigit)
}
