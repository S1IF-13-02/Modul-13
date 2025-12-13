package main

import "fmt"

func main() {
	var kata string
	var jumlah, i int
	var stop bool
	fmt.Scan(&kata, &jumlah)

	i = 0

	for stop = false; !stop; {
		fmt.Println(kata)
		i++
		stop = i >= jumlah
	}
}
