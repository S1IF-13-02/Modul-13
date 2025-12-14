package main

import "fmt"

func main() {
	var x, y int
	var isKelipatan bool

	fmt.Print("Masukkan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scanln(&y)

	for {
		x -= y
		fmt.Println(x)
		if x == 0 {
			isKelipatan = true
			break
		}

		if x < 0 {
			isKelipatan = false
			break
		}
	}
	fmt.Println(isKelipatan)
}