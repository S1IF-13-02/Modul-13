package main

import "fmt"

func main() {
	var n int
	var JumlahDigit int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)
	JumlahDigit = 0
	for {
		n = n / 10
		JumlahDigit++

		if n == 0 {
			break
		}
	}
	fmt.Print("Jumlah digit: ", JumlahDigit)
}
