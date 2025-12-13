package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	hasil := x
	for {
		hasil = hasil - y
		fmt.Println(hasil)

		if hasil == 0 {
			fmt.Println("true")
			break
		}

		if hasil < 0 {
			fmt.Println("false")
			break
		}
	}
}
