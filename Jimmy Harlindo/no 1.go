package main

import "fmt"


func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	hitung := 0

	for {
		hitung++
		n = n / 10
		if n == 0 {
			break
		}
	}
	fmt.Println("Jumlah digit:", hitung)
}
