package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&a)

	hitung := 0
	for selesai := false; !selesai; {
		a = a / 10
		hitung++
		selesai = a == 0
	}
	fmt.Println(hitung)
}
