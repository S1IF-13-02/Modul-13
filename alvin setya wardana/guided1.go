package main

import (
	"fmt"
)

func main() {
	var kata string
	var jumlah int

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)

	fmt.Print("Masukkan jumlah pengulangan: ")
	fmt.Scanln(&jumlah)

	i := 0
	for {
		fmt.Println(kata)
		i++
		if i >= jumlah {
			break
		}
	}
}
