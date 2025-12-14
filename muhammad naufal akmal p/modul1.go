package main

import "fmt"

func main() {
	var kata string
	var jumlah int

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	fmt.Print("Masukkan jumlah pengulangan: ")
	fmt.Scan(&jumlah)

	counter := 0
	for {
		fmt.Println(kata)
		counter++
		if counter >= jumlah {
			break
		}
	}
	fmt.Println("Selesai.")
}