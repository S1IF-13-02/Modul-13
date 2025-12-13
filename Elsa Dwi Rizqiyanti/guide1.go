package main

import "fmt"

func main() {
	var kata string
	var jumlah int

	fmt.Scan(&kata, &jumlah)

	counter := 0

	for {
		fmt.Println(kata)
		counter++

		if counter == jumlah {
			break
		}
	}
}
