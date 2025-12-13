package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan bilangan : ")
	fmt.Scan(&n)

	jumlahDigit := 0

	for {
		jumlahDigit++
		n = n / 10

		if n == 0 {
			break
		}
	}

	fmt.Println(jumlahDigit)
}
