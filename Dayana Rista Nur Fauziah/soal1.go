package main

import "fmt"

func main() {
	var a int
	jumlah := 0

	fmt.Scan(&a)

	for a > 0 {
		a = a / 10
		jumlah++
	}

	fmt.Println("Jumlah digit:", jumlah)
}
