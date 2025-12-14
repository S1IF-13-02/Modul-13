package main

import "fmt" 

func main() {
	var n int

	for { 
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&n)

		if n > 0 {
		fmt.Print(n, " adalah bilangan bulat positif")
		break
		}
	}
}
