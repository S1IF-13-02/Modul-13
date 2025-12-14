package main

import "fmt"

func main() {
	var kata string
	var jumlah int

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	fmt.Print("Masukkan jumlah kata yang ingin diulang: ")
	fmt.Scan(&jumlah)

	i := 0
	for i < jumlah {
		fmt.Println(kata)
		i++
	}
}
