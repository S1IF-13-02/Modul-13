package main

import "fmt"

func main() {
	var bil float64
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&bil)

	i := int(bil * 10)
	end := int(bil)*10 + 10

	for {
		i++
		fmt.Printf("%.1f\n", float64(i)/10)

		if i >= end {
			break
		}
	}
}
