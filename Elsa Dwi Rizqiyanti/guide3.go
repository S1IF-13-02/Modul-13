package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	sisa := x

	for {
		sisa = sisa - y
		fmt.Println(sisa)

		if sisa <= 0 {
			break
		}
	}
	if sisa == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
