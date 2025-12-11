package main

import (
	"fmt"
)

func main() {
	var kata string
	var jumlah int

	fmt.Scan(&kata, &jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Println(kata)
	}
}