package main

import "fmt"

func main() {
	var angka int

	for{
		fmt.Scan(&angka)

		if angka > 0 {
			break
		}
	}
	fmt.Println(angka, "adalah bilangan bulat positif")
}