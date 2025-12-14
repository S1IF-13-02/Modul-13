package main

import "fmt"

func main() {
	var x, y int
	var isKelipatan bool

	fmt.Scan(&x)
	fmt.Scan(&y)

	sisa := x
	isKelipatan = true

	for {
		sisa = sisa - y
		fmt.Println(sisa)

		if sisa < 0 {
			isKelipatan = false
			break
		}

		if sisa == 0 {
			break
		}
	}

	fmt.Println(isKelipatan)
}
