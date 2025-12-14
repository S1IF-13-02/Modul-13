package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	jumlahDigit := 0

	if n == 0 {
		jumlahDigit = 1
	} else {
		for n > 0 {
			jumlahDigit++ 
			n = n / 10
		}
	}
	fmt.Println(jumlahDigit)
}