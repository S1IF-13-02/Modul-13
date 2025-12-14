package main

import "fmt"

func main() {
	var kata string
	var perulangan int
	var hitung int = 0

	fmt.Scan(&kata)
	fmt.Scan(&perulangan)

	for hitung < perulangan {
		fmt.Println(kata)
		hitung++
	}
}
