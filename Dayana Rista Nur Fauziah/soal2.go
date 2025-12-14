package main

import "fmt"

func main() {
	var x float64
	var nilai int
	var batas int

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)

	nilai = int(x * 10)
	batas = (int(x) + 1) * 10

	for nilai < batas {
		nilai = nilai + 1
		fmt.Println(float64(nilai) / 10)
	}
}
