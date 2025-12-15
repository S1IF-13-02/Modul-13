package main

import "fmt"

func main() {
	var x, y int
	var hasil bool

	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	for {
		x -= y
		fmt.Println(x)
		if x == 0 {
			hasil = true
			break
		}
		if x < 0 {
			hasil = false
			break
		}

	}
	fmt.Println(hasil)
}
