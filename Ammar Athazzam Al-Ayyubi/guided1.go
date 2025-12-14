package main

import "fmt"

func main() {
	var kata string
	var n int

	fmt.Print("Masukkan sebuah kata dan jumlah Perulangannya: ")
	fmt.Scanln(&kata, &n)

	for {
		if n <= 0 {
			break
		}
		fmt.Println(kata)
		n--
	}
}
