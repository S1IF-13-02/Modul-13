package main

import "fmt"

func main() {
	var kata string
	var loop int
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	fmt.Print("Masukkan loop: ")
	fmt.Scan(&loop)
	hitung := 0
	for stop := false; !stop; {
		fmt.Println(kata)
		hitung++
		stop = hitung == loop

	}
}
