package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for n > 0 {
		n = n / 10
		jumlah++
	}
	fmt.Println(jumlah)
}
