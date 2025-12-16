package main

import "fmt"

func main() {
	var n float64
	var nilai int
	var batas int
	var selesai bool

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&n)

	nilai = int(n * 10)

	if nilai%10 == 0 {
		batas = nilai
	} else {
		batas = (nilai/10 + 1) * 10
	}

	for selesai = false; !selesai; {
		nilai++
		fmt.Printf("%.1f\n", float64(nilai)/10)
		selesai = nilai == batas
	}
}
