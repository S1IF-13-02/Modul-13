package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	for {
		x = x - y
		fmt.Println(x)

		if x <= 0 {
			break
		}
	}

	fmt.Println(x == 0)
}
