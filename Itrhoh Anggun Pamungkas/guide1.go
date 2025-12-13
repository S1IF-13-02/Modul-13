package main

import "fmt"

func main() {
	var kata string
	var jml int

	fmt.Scan(&kata, &jml)

	hitung:=0

	for{
		fmt.Println(kata)
		hitung++
	
		if hitung == jml{
			break
		}
	}
}